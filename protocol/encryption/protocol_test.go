package encryption

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/status-im/status-go/protocol/tt"

	"github.com/status-im/status-go/protocol/sqlite"

	"github.com/status-im/status-go/eth-node/crypto"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/status-im/status-go/protocol/encryption/multidevice"
	"github.com/status-im/status-go/protocol/encryption/sharedsecret"
)

func TestProtocolServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProtocolServiceTestSuite))
}

type ProtocolServiceTestSuite struct {
	suite.Suite
	aliceDBPath *os.File
	bobDBPath   *os.File
	alice       *Protocol
	bob         *Protocol
	logger      *zap.Logger
}

func (s *ProtocolServiceTestSuite) SetupTest() {
	var err error

	s.logger = tt.MustCreateTestLogger()

	s.aliceDBPath, err = ioutil.TempFile("", "alice.db.sql")
	s.Require().NoError(err)
	aliceDBKey := "alice"

	s.bobDBPath, err = ioutil.TempFile("", "bob.db.sql")
	s.Require().NoError(err)
	bobDBKey := "bob"

	addedBundlesHandler := func(addedBundles []*multidevice.Installation) {}
	onNewSharedSecretHandler := func(secret []*sharedsecret.Secret) {}

	db, err := sqlite.Open(s.aliceDBPath.Name(), aliceDBKey)
	s.Require().NoError(err)
	s.alice = New(
		db,
		"1",
		addedBundlesHandler,
		onNewSharedSecretHandler,
		func(*ProtocolMessageSpec) {},
		s.logger.With(zap.String("user", "alice")),
	)

	db, err = sqlite.Open(s.bobDBPath.Name(), bobDBKey)
	s.Require().NoError(err)
	s.bob = New(
		db,
		"2",
		addedBundlesHandler,
		onNewSharedSecretHandler,
		func(*ProtocolMessageSpec) {},
		s.logger.With(zap.String("user", "bob")),
	)
}

func (s *ProtocolServiceTestSuite) TearDownTest() {
	os.Remove(s.aliceDBPath.Name())
	os.Remove(s.bobDBPath.Name())
	_ = s.logger.Sync()
}

func (s *ProtocolServiceTestSuite) TestBuildPublicMessage() {
	aliceKey, err := crypto.GenerateKey()
	s.NoError(err)

	payload := []byte("test")
	s.NoError(err)

	msg, err := s.alice.BuildPublicMessage(aliceKey, payload)
	s.NoError(err)
	s.NotNil(msg, "It creates a message")

	s.NotNilf(msg.Message.GetBundles(), "It adds a bundle to the message")
}

func (s *ProtocolServiceTestSuite) TestBuildDirectMessage() {
	bobKey, err := crypto.GenerateKey()
	s.NoError(err)
	aliceKey, err := crypto.GenerateKey()
	s.NoError(err)

	payload := []byte("test")

	msgSpec, err := s.alice.BuildDirectMessage(aliceKey, &bobKey.PublicKey, payload)
	s.NoError(err)
	s.NotNil(msgSpec, "It creates a message spec")

	msg := msgSpec.Message
	s.NotNil(msg, "It creates a messages")

	s.NotNilf(msg.GetBundles(), "It adds a bundle to the message")

	directMessage := msg.GetDirectMessage()
	s.NotNilf(directMessage, "It sets the direct message")

	encryptedPayload := directMessage["none"].GetPayload()
	s.NotNilf(encryptedPayload, "It sets the payload of the message")

	s.NotEqualf(payload, encryptedPayload, "It encrypts the payload")
}

func (s *ProtocolServiceTestSuite) TestBuildAndReadDirectMessage() {
	bobKey, err := crypto.GenerateKey()
	s.Require().NoError(err)
	aliceKey, err := crypto.GenerateKey()
	s.Require().NoError(err)

	payload := []byte("test")

	// Message is sent with DH
	msgSpec, err := s.alice.BuildDirectMessage(aliceKey, &bobKey.PublicKey, payload)
	s.Require().NoError(err)
	s.Require().NotNil(msgSpec)

	msg := msgSpec.Message
	s.Require().NotNil(msg)

	// Bob is able to decrypt the message
	unmarshaledMsg, err := s.bob.HandleMessage(bobKey, &aliceKey.PublicKey, msg, []byte("message-id"))
	s.NoError(err)
	s.NotNil(unmarshaledMsg)

	recoveredPayload := []byte("test")
	s.Equalf(payload, recoveredPayload, "It successfully unmarshal the decrypted message")
}

func (s *ProtocolServiceTestSuite) TestSecretNegotiation() {
	var secretResponse []*sharedsecret.Secret
	bobKey, err := crypto.GenerateKey()
	s.NoError(err)
	aliceKey, err := crypto.GenerateKey()
	s.NoError(err)

	payload := []byte("test")

	s.bob.onNewSharedSecretHandler = func(secret []*sharedsecret.Secret) {
		secretResponse = secret
	}
	msgSpec, err := s.alice.BuildDirectMessage(aliceKey, &bobKey.PublicKey, payload)
	s.NoError(err)
	s.NotNil(msgSpec, "It creates a message spec")

	bundle := msgSpec.Message.GetBundles()[0]
	s.Require().NotNil(bundle)

	signedPreKeys := bundle.GetSignedPreKeys()
	s.Require().NotNil(signedPreKeys)

	signedPreKey := signedPreKeys["1"]
	s.Require().NotNil(signedPreKey)

	s.Require().Equal(uint32(1), signedPreKey.GetProtocolVersion())

	_, err = s.bob.HandleMessage(bobKey, &aliceKey.PublicKey, msgSpec.Message, []byte("message-id"))
	s.NoError(err)

	s.Require().NotNil(secretResponse)
}

func (s *ProtocolServiceTestSuite) TestPropagatingSavedSharedSecretsOnStart() {
	var secretResponse []*sharedsecret.Secret

	aliceKey, err := crypto.GenerateKey()
	s.NoError(err)
	bobKey, err := crypto.GenerateKey()
	s.NoError(err)

	// Generate and save a shared secret.
	generatedSecret, err := s.alice.secret.Generate(aliceKey, &bobKey.PublicKey, "installation-1")
	s.NoError(err)

	s.alice.onNewSharedSecretHandler = func(secret []*sharedsecret.Secret) {
		secretResponse = secret
	}

	err = s.alice.Start(aliceKey)
	s.NoError(err)

	s.Require().NotNil(secretResponse)
	s.Require().Len(secretResponse, 1)
	s.Equal(crypto.FromECDSAPub(generatedSecret.Identity), crypto.FromECDSAPub(secretResponse[0].Identity))
	s.Equal(generatedSecret.Key, secretResponse[0].Key)
}
