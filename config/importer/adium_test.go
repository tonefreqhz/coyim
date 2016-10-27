package importer

import (
	"github.com/twstrike/coyim/config"
	. "gopkg.in/check.v1"
)

type AdiumSuite struct{}

var _ = Suite(&AdiumSuite{})

func (s *AdiumSuite) Test_AdiumImporter_canImportAccountMappings(c *C) {
	importer := adiumImporter{}
	mappings, ok := importer.readAccountMappings(testResourceFilename("adium_test_data/Accounts.plist"))

	c.Assert(ok, Equals, true)
	c.Assert(len(mappings), Equals, 6)
	c.Assert(mappings["2"], Equals, adiumAccountMapping{objectID: "2", uid: "foobaar@mesopotamia.xmpp", accountType: "libpurple-jabber-gtalk"})
	c.Assert(mappings["5"], Equals, adiumAccountMapping{objectID: "5", uid: "baarfoo@coyim.com", accountType: "libpurple-Jabber"})
	c.Assert(mappings["10"], Equals, adiumAccountMapping{objectID: "10", uid: "baarfoo@gmail.com", accountType: "libpurple-jabber-gtalk"})
	c.Assert(mappings["11"], Equals, adiumAccountMapping{objectID: "11", uid: "strike-test@dukgo.com", accountType: "libpurple-Jabber"})
	c.Assert(mappings["12"], Equals, adiumAccountMapping{objectID: "12", uid: "strike-test@jabber.otr.im", accountType: "libpurple-Jabber"})
	c.Assert(mappings["13"], Equals, adiumAccountMapping{objectID: "13", uid: "strike-test@jabber.calyxinstitute.org", accountType: "libpurple-Jabber"})
}

func (s *AdiumSuite) Test_AdiumImporter_canImportKeysFromFile(c *C) {
	importer := adiumImporter{}
	res, ok := importer.importKeysFrom(testResourceFilename("adium_test_data/otr.private_key"))
	c.Assert(ok, Equals, true)
	c.Assert(len(res), Equals, 3)
	c.Assert(res["2"], DeepEquals, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x87, 0xc1, 0x7f, 0x67, 0x46, 0x72, 0x28, 0xe1, 0x78, 0x33, 0x5a, 0xa7, 0xe8, 0x9e, 0x68, 0x75, 0xac, 0xc, 0x29, 0x2d, 0xb, 0x6b, 0xfe, 0x40, 0xc9, 0xbd, 0x16, 0xe3, 0x5d, 0x0, 0x11, 0x5, 0xd2, 0xd0, 0x59, 0x48, 0xce, 0x99, 0xcf, 0xaa, 0x48, 0xf6, 0xda, 0x7e, 0xe3, 0xaf, 0x82, 0x8d, 0xc3, 0x1f, 0xf8, 0xc2, 0xfb, 0x93, 0xbf, 0xb7, 0xb2, 0xfe, 0x76, 0x86, 0x2, 0x39, 0xfd, 0x2, 0x74, 0xf7, 0xdb, 0x95, 0x46, 0x0, 0x8f, 0xde, 0x2e, 0x97, 0xeb, 0x90, 0x2a, 0xe, 0xad, 0x9, 0x97, 0x2d, 0x3b, 0x1a, 0x5, 0x37, 0xb1, 0x43, 0x80, 0xa4, 0x74, 0x46, 0x7, 0x83, 0xa, 0x99, 0xf1, 0x0, 0xcc, 0x36, 0xb, 0xd8, 0x16, 0x9c, 0xce, 0x9c, 0x19, 0x62, 0x7a, 0x31, 0x27, 0xec, 0xbc, 0xf, 0xdb, 0x50, 0xd4, 0xf8, 0xe9, 0x75, 0x50, 0x69, 0xe2, 0xb, 0x82, 0x82, 0x3, 0x3, 0x0, 0x0, 0x0, 0x14, 0x8d, 0x32, 0xdb, 0xfd, 0x90, 0xde, 0x65, 0x7a, 0xaf, 0xd1, 0x4f, 0xfc, 0xd3, 0xb2, 0x1a, 0x7f, 0xa3, 0x98, 0x45, 0x49, 0x0, 0x0, 0x0, 0x80, 0x13, 0x1c, 0xd5, 0xa2, 0xe1, 0x9c, 0x1e, 0xea, 0x82, 0xb5, 0xad, 0x6e, 0x5d, 0x9c, 0x63, 0x52, 0x58, 0x17, 0xc3, 0xb3, 0x99, 0x50, 0xac, 0x1f, 0x4b, 0x4a, 0x1c, 0x1e, 0xee, 0xd0, 0x9a, 0xe9, 0x5d, 0x6, 0xf6, 0x3a, 0x57, 0x19, 0x95, 0xf9, 0xb9, 0xff, 0x4e, 0x7, 0xe3, 0xfc, 0xdd, 0xc0, 0xfc, 0x97, 0xee, 0x88, 0xa5, 0xf6, 0x48, 0xa9, 0x30, 0x80, 0x5e, 0xf7, 0x34, 0xf4, 0xed, 0x29, 0xe7, 0x18, 0xaf, 0x93, 0x9a, 0x76, 0x6b, 0xc5, 0x4b, 0x5f, 0x9b, 0x43, 0xce, 0x3e, 0x70, 0x33, 0x99, 0xd7, 0xb1, 0xa6, 0x8e, 0x4b, 0x7c, 0xb0, 0x23, 0x9a, 0x42, 0xee, 0x2c, 0x68, 0xb0, 0x6f, 0xe2, 0xb5, 0xab, 0x59, 0xf7, 0xa9, 0x26, 0xaf, 0x96, 0xed, 0xaa, 0xe6, 0x86, 0x95, 0x43, 0x78, 0x63, 0xe7, 0x6e, 0xa6, 0x90, 0x39, 0xcd, 0x76, 0x92, 0xa, 0x83, 0x7b, 0xc4, 0x6f, 0x1b, 0x38, 0x0, 0x0, 0x0, 0x80, 0x3, 0x85, 0xe6, 0xcc, 0x5, 0xe5, 0x1b, 0x4a, 0x3f, 0x45, 0xdd, 0xc8, 0x58, 0xec, 0x4c, 0x77, 0x9, 0x99, 0x47, 0xf1, 0x88, 0x8b, 0x6e, 0xe4, 0x26, 0xf3, 0xc4, 0x35, 0x69, 0xbd, 0xf2, 0xc, 0xbb, 0xa6, 0xe5, 0x50, 0x6, 0xec, 0xdd, 0x98, 0xf4, 0x53, 0xfa, 0x20, 0xf0, 0x6c, 0x38, 0xe3, 0xf3, 0x39, 0x6d, 0x8a, 0x3f, 0x40, 0xea, 0x50, 0xac, 0xd9, 0x59, 0x30, 0xa3, 0xb4, 0xf8, 0xf2, 0x79, 0xb8, 0x65, 0x56, 0xca, 0xab, 0x7e, 0xe6, 0xdc, 0x55, 0xe1, 0x76, 0x3e, 0x2f, 0x5d, 0x36, 0x50, 0xc, 0xf6, 0x72, 0xef, 0x94, 0x14, 0x96, 0x32, 0xdc, 0x49, 0x91, 0xc0, 0x25, 0x86, 0x88, 0x7a, 0x39, 0x67, 0x27, 0x46, 0x0, 0x40, 0x3c, 0xb8, 0x78, 0x90, 0xfc, 0x9, 0x69, 0x8a, 0x47, 0xf2, 0x50, 0xb2, 0x1d, 0xbe, 0x46, 0x62, 0x44, 0x58, 0x21, 0x6a, 0xe9, 0x5a, 0x6, 0x47, 0x11, 0x0, 0x0, 0x0, 0x14, 0x29, 0x41, 0x1a, 0x59, 0x74, 0x58, 0x75, 0xf, 0x7a, 0x5f, 0x2b, 0xd5, 0x61, 0x85, 0xdf, 0x71, 0x93, 0xf, 0xd4, 0x2e})
	c.Assert(res["5"], DeepEquals, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x87, 0xc1, 0x7f, 0x67, 0x46, 0x72, 0x28, 0xe1, 0x78, 0x33, 0x5a, 0xa7, 0xe8, 0x9e, 0x68, 0x75, 0xac, 0xc, 0x29, 0x2d, 0xb, 0x6b, 0xfe, 0x40, 0xc9, 0xbd, 0x16, 0xe3, 0x5d, 0x0, 0x11, 0x5, 0xd2, 0xd0, 0x59, 0x48, 0xce, 0x99, 0xcf, 0xaa, 0x48, 0xf6, 0xda, 0x7e, 0xe3, 0xaf, 0x82, 0x8d, 0xc3, 0x1f, 0xf8, 0xc2, 0xfb, 0x93, 0xbf, 0xb7, 0xb2, 0xfe, 0x76, 0x86, 0x2, 0x39, 0xfd, 0x2, 0x74, 0xf7, 0xdb, 0x95, 0x46, 0x0, 0x8f, 0xde, 0x2e, 0x97, 0xeb, 0x90, 0x2a, 0xe, 0xad, 0x9, 0x97, 0x2d, 0x3b, 0x1a, 0x5, 0x37, 0xb1, 0x43, 0x80, 0xa4, 0x74, 0x46, 0x7, 0x83, 0xa, 0x99, 0xf1, 0x0, 0xcc, 0x36, 0xb, 0xd8, 0x16, 0x9c, 0xce, 0x9c, 0x19, 0x62, 0x7a, 0x31, 0x27, 0xec, 0xbc, 0xf, 0xdb, 0x50, 0xd4, 0xf8, 0xe9, 0x75, 0x50, 0x69, 0xe2, 0xb, 0x82, 0x82, 0x3, 0x3, 0x0, 0x0, 0x0, 0x14, 0x8d, 0x32, 0xdb, 0xfd, 0x90, 0xde, 0x65, 0x7a, 0xaf, 0xd1, 0x4f, 0xfc, 0xd3, 0xb2, 0x1a, 0x7f, 0xa3, 0x98, 0x45, 0x49, 0x0, 0x0, 0x0, 0x80, 0x13, 0x1c, 0xd5, 0xa2, 0xe1, 0x9c, 0x1e, 0xea, 0x82, 0xb5, 0xad, 0x6e, 0x5d, 0x9c, 0x63, 0x52, 0x58, 0x17, 0xc3, 0xb3, 0x99, 0x50, 0xac, 0x1f, 0x4b, 0x4a, 0x1c, 0x1e, 0xee, 0xd0, 0x9a, 0xe9, 0x5d, 0x6, 0xf6, 0x3a, 0x57, 0x19, 0x95, 0xf9, 0xb9, 0xff, 0x4e, 0x7, 0xe3, 0xfc, 0xdd, 0xc0, 0xfc, 0x97, 0xee, 0x88, 0xa5, 0xf6, 0x48, 0xa9, 0x30, 0x80, 0x5e, 0xf7, 0x34, 0xf4, 0xed, 0x29, 0xe7, 0x18, 0xaf, 0x93, 0x9a, 0x76, 0x6b, 0xc5, 0x4b, 0x5f, 0x9b, 0x43, 0xce, 0x3e, 0x70, 0x33, 0x99, 0xd7, 0xb1, 0xa6, 0x8e, 0x4b, 0x7c, 0xb0, 0x23, 0x9a, 0x42, 0xee, 0x2c, 0x68, 0xb0, 0x6f, 0xe2, 0xb5, 0xab, 0x59, 0xf7, 0xa9, 0x26, 0xaf, 0x96, 0xed, 0xaa, 0xe6, 0x86, 0x95, 0x43, 0x78, 0x63, 0xe7, 0x6e, 0xa6, 0x90, 0x39, 0xcd, 0x76, 0x92, 0xa, 0x83, 0x7b, 0xc4, 0x6f, 0x1b, 0x38, 0x0, 0x0, 0x0, 0x80, 0x3, 0x85, 0xe6, 0xcc, 0x5, 0xe5, 0x1b, 0x4a, 0x3f, 0x45, 0xdd, 0xc8, 0x58, 0xec, 0x4c, 0x77, 0x9, 0x99, 0x47, 0xf1, 0x88, 0x8b, 0x6e, 0xe4, 0x26, 0xf3, 0xc4, 0x35, 0x69, 0xbd, 0xf2, 0xc, 0xbb, 0xa6, 0xe5, 0x50, 0x6, 0xec, 0xdd, 0x98, 0xf4, 0x53, 0xfa, 0x20, 0xf0, 0x6c, 0x38, 0xe3, 0xf3, 0x39, 0x6d, 0x8a, 0x3f, 0x40, 0xea, 0x50, 0xac, 0xd9, 0x59, 0x30, 0xa3, 0xb4, 0xf8, 0xf2, 0x79, 0xb8, 0x65, 0x56, 0xca, 0xab, 0x7e, 0xe6, 0xdc, 0x55, 0xe1, 0x76, 0x3e, 0x2f, 0x5d, 0x36, 0x50, 0xc, 0xf6, 0x72, 0xef, 0x94, 0x14, 0x96, 0x32, 0xdc, 0x49, 0x91, 0xc0, 0x25, 0x86, 0x88, 0x7a, 0x39, 0x67, 0x27, 0x46, 0x0, 0x40, 0x3c, 0xb8, 0x78, 0x90, 0xfc, 0x9, 0x69, 0x8a, 0x47, 0xf2, 0x50, 0xb2, 0x1d, 0xbe, 0x46, 0x62, 0x44, 0x58, 0x21, 0x6a, 0xe9, 0x5a, 0x6, 0x47, 0x11, 0x0, 0x0, 0x0, 0x14, 0x39, 0x41, 0x1a, 0x59, 0x74, 0x58, 0x75, 0xf, 0x7a, 0x5f, 0x2b, 0xd5, 0x61, 0x85, 0xdf, 0x71, 0x93, 0xf, 0xd4, 0x2e})
	c.Assert(res["10"], DeepEquals, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x87, 0xc1, 0x7f, 0x67, 0x46, 0x72, 0x28, 0xe1, 0x78, 0x33, 0x5a, 0xa7, 0xe8, 0x9e, 0x68, 0x75, 0xac, 0xc, 0x29, 0x2d, 0xb, 0x6b, 0xfe, 0x40, 0xc9, 0xbd, 0x16, 0xe3, 0x5d, 0x0, 0x11, 0x5, 0xd2, 0xd0, 0x59, 0x48, 0xce, 0x99, 0xcf, 0xaa, 0x48, 0xf6, 0xda, 0x7e, 0xe3, 0xaf, 0x82, 0x8d, 0xc3, 0x1f, 0xf8, 0xc2, 0xfb, 0x93, 0xbf, 0xb7, 0xb2, 0xfe, 0x76, 0x86, 0x2, 0x39, 0xfd, 0x2, 0x74, 0xf7, 0xdb, 0x95, 0x46, 0x0, 0x8f, 0xde, 0x2e, 0x97, 0xeb, 0x90, 0x2a, 0xe, 0xad, 0x9, 0x97, 0x2d, 0x3b, 0x1a, 0x5, 0x37, 0xb1, 0x43, 0x80, 0xa4, 0x74, 0x46, 0x7, 0x83, 0xa, 0x99, 0xf1, 0x0, 0xcc, 0x36, 0xb, 0xd8, 0x16, 0x9c, 0xce, 0x9c, 0x19, 0x62, 0x7a, 0x31, 0x27, 0xec, 0xbc, 0xf, 0xdb, 0x50, 0xd4, 0xf8, 0xe9, 0x75, 0x50, 0x69, 0xe2, 0xb, 0x82, 0x82, 0x3, 0x3, 0x0, 0x0, 0x0, 0x14, 0x8d, 0x32, 0xdb, 0xfd, 0x90, 0xde, 0x65, 0x7a, 0xaf, 0xd1, 0x4f, 0xfc, 0xd3, 0xb2, 0x1a, 0x7f, 0xa3, 0x98, 0x45, 0x49, 0x0, 0x0, 0x0, 0x80, 0x13, 0x1c, 0xd5, 0xa2, 0xe1, 0x9c, 0x1e, 0xea, 0x82, 0xb5, 0xad, 0x6e, 0x5d, 0x9c, 0x63, 0x52, 0x58, 0x17, 0xc3, 0xb3, 0x99, 0x50, 0xac, 0x1f, 0x4b, 0x4a, 0x1c, 0x1e, 0xee, 0xd0, 0x9a, 0xe9, 0x5d, 0x6, 0xf6, 0x3a, 0x57, 0x19, 0x95, 0xf9, 0xb9, 0xff, 0x4e, 0x7, 0xe3, 0xfc, 0xdd, 0xc0, 0xfc, 0x97, 0xee, 0x88, 0xa5, 0xf6, 0x48, 0xa9, 0x30, 0x80, 0x5e, 0xf7, 0x34, 0xf4, 0xed, 0x29, 0xe7, 0x18, 0xaf, 0x93, 0x9a, 0x76, 0x6b, 0xc5, 0x4b, 0x5f, 0x9b, 0x43, 0xce, 0x3e, 0x70, 0x33, 0x99, 0xd7, 0xb1, 0xa6, 0x8e, 0x4b, 0x7c, 0xb0, 0x23, 0x9a, 0x42, 0xee, 0x2c, 0x68, 0xb0, 0x6f, 0xe2, 0xb5, 0xab, 0x59, 0xf7, 0xa9, 0x26, 0xaf, 0x96, 0xed, 0xaa, 0xe6, 0x86, 0x95, 0x43, 0x78, 0x63, 0xe7, 0x6e, 0xa6, 0x90, 0x39, 0xcd, 0x76, 0x92, 0xa, 0x83, 0x7b, 0xc4, 0x6f, 0x1b, 0x38, 0x0, 0x0, 0x0, 0x80, 0x3, 0x85, 0xe6, 0xcc, 0x5, 0xe5, 0x1b, 0x4a, 0x3f, 0x45, 0xdd, 0xc8, 0x58, 0xec, 0x4c, 0x77, 0x9, 0x99, 0x47, 0xf1, 0x88, 0x8b, 0x6e, 0xe4, 0x26, 0xf3, 0xc4, 0x35, 0x69, 0xbd, 0xf2, 0xc, 0xbb, 0xa6, 0xe5, 0x50, 0x6, 0xec, 0xdd, 0x98, 0xf4, 0x53, 0xfa, 0x20, 0xf0, 0x6c, 0x38, 0xe3, 0xf3, 0x39, 0x6d, 0x8a, 0x3f, 0x40, 0xea, 0x50, 0xac, 0xd9, 0x59, 0x30, 0xa3, 0xb4, 0xf8, 0xf2, 0x79, 0xb8, 0x65, 0x56, 0xca, 0xab, 0x7e, 0xe6, 0xdc, 0x55, 0xe1, 0x76, 0x3e, 0x2f, 0x5d, 0x36, 0x50, 0xc, 0xf6, 0x72, 0xef, 0x94, 0x14, 0x96, 0x32, 0xdc, 0x49, 0x91, 0xc0, 0x25, 0x86, 0x88, 0x7a, 0x39, 0x67, 0x27, 0x46, 0x0, 0x40, 0x3c, 0xb8, 0x78, 0x90, 0xfc, 0x9, 0x69, 0x8a, 0x47, 0xf2, 0x50, 0xb2, 0x1d, 0xbe, 0x46, 0x62, 0x44, 0x58, 0x21, 0x6a, 0xe9, 0x5a, 0x6, 0x47, 0x11, 0x0, 0x0, 0x0, 0x14, 0x49, 0x41, 0x1a, 0x59, 0x74, 0x58, 0x75, 0xf, 0x7a, 0x5f, 0x2b, 0xd5, 0x61, 0x85, 0xdf, 0x71, 0x93, 0xf, 0xd4, 0x2e})
}

func (s *AdiumSuite) Test_AdiumImporter_canImportFingerprintsFromFile(c *C) {
	importer := adiumImporter{}

	res, ok := importer.importFingerprintsFrom(testResourceFilename("adium_test_data/otr.fingerprints"))

	c.Assert(ok, Equals, true)
	c.Assert(len(res), Equals, 2)

	c.Assert(len(res["2"]), Equals, 5)
	c.Assert(len(res["5"]), Equals, 5)

	c.Check(res["2"][0].UserID, Equals, "abdfdergsfdfgd@gmail.com")
	c.Check(res["2"][0].Fingerprint, DeepEquals, decode("1bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["2"][0].Untrusted, Equals, false)

	c.Check(res["2"][1].UserID, Equals, "llaall@gmail.com")
	c.Check(res["2"][1].Fingerprint, DeepEquals, decode("7bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["2"][1].Untrusted, Equals, false)

	c.Check(res["2"][2].UserID, Equals, "abc@jabber.org")
	c.Check(res["2"][2].Fingerprint, DeepEquals, decode("8bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["2"][2].Untrusted, Equals, false)

	c.Check(res["2"][3].UserID, Equals, "xxqqx@coy.im")
	c.Check(res["2"][3].Fingerprint, DeepEquals, decode("9bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["2"][3].Untrusted, Equals, false)

	c.Check(res["2"][4].UserID, Equals, "xxqqx@coy.im")
	c.Check(res["2"][4].Fingerprint, DeepEquals, decode("10cdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["2"][4].Untrusted, Equals, false)

	c.Check(res["5"][0].UserID, Equals, "hmmabc@mesopotamia.xmpp")
	c.Check(res["5"][0].Fingerprint, DeepEquals, decode("2bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["5"][0].Untrusted, Equals, false)

	c.Check(res["5"][1].UserID, Equals, "hmmabc@mesopotamia.xmpp")
	c.Check(res["5"][1].Fingerprint, DeepEquals, decode("3bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["5"][1].Untrusted, Equals, true)

	c.Check(res["5"][2].UserID, Equals, "abcaaaa@mesopotamia.xmpp")
	c.Check(res["5"][2].Fingerprint, DeepEquals, decode("4bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["5"][2].Untrusted, Equals, false)

	c.Check(res["5"][3].UserID, Equals, "fooba@mesopotamia.xmpp")
	c.Check(res["5"][3].Fingerprint, DeepEquals, decode("5bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["5"][3].Untrusted, Equals, true)

	c.Check(res["5"][4].UserID, Equals, "fooba@mesopotamia.xmpp")
	c.Check(res["5"][4].Fingerprint, DeepEquals, decode("6bcdefabcdefabcdefabcdefabcdefabcdefcdef"))
	c.Check(res["5"][4].Untrusted, Equals, false)
}

func (s *AdiumSuite) Test_AdiumImporter_canImportAccountsFromFile(c *C) {
	importer := adiumImporter{}
	res, ok := importer.importAccounts(testResourceFilename("adium_test_data/libpurple/accounts.xml"))

	c.Assert(ok, Equals, true)
	c.Assert(len(res), Equals, 3)

	c.Check(res["baarfoo@coyim.com"].Account, Equals, "baarfoo@coyim.com")
	c.Check(res["baarfoo@coyim.com"].Server, Equals, "talk.google.com")
	c.Check(res["baarfoo@coyim.com"].Password, Equals, "")
	c.Check(res["baarfoo@coyim.com"].Port, Equals, 5222)
	c.Check(len(res["baarfoo@coyim.com"].Proxies), Equals, 0)

	c.Check(res["baarfoo@coyim.org"].Account, Equals, "baarfoo@coyim.org")
	c.Check(res["baarfoo@coyim.org"].Server, Equals, "")
	c.Check(res["baarfoo@coyim.org"].Password, Equals, "")
	c.Check(res["baarfoo@coyim.org"].Port, Equals, 5222)
	c.Check(len(res["baarfoo@coyim.com"].Proxies), Equals, 0)

	c.Check(res["foobaar@mesopotamia.xmpp"].Account, Equals, "foobaar@mesopotamia.xmpp")
	c.Check(res["foobaar@mesopotamia.xmpp"].Server, Equals, "talk.google.com")
	c.Check(res["foobaar@mesopotamia.xmpp"].Password, Equals, "")
	c.Check(res["foobaar@mesopotamia.xmpp"].Port, Equals, 5222)
	c.Check(len(res["foobaar@mesopotamia.xmpp"].Proxies), Equals, 0)
}

func (s *AdiumSuite) Test_AdiumImporter_canImportGlobalOTRPrefs(c *C) {
	importer := adiumImporter{}
	_, ok := importer.importGlobalPrefs(testResourceFilename("adium_test_data/libpurple/prefs.xml"))
	c.Assert(ok, Equals, false)
}

func (s *AdiumSuite) Test_AdiumImporter_canImportBuddyOTRPrefs(c *C) {
	importer := adiumImporter{}
	res, ok := importer.importPeerPrefs(testResourceFilename("adium_test_data/libpurple/blist.xml"))

	c.Assert(ok, Equals, true)
	c.Assert(len(res), Equals, 1)
	c.Assert(len(res["baarfoo@coyim.com"]), Equals, 1)
	c.Assert(res["baarfoo@coyim.com"]["abcaaaa@mesopotamia.xmpp"].enabled, Equals, true)
	c.Assert(res["baarfoo@coyim.com"]["abcaaaa@mesopotamia.xmpp"].automatic, Equals, false)
	c.Assert(res["baarfoo@coyim.com"]["abcaaaa@mesopotamia.xmpp"].avoidLoggingOTR, Equals, true)
	c.Assert(res["baarfoo@coyim.com"]["abcaaaa@mesopotamia.xmpp"].onlyPrivate, Equals, false)
}

func (s *AdiumSuite) Test_AdiumImporter_canDoAFullImport(c *C) {
	importer := adiumImporter{}
	res, ok := importer.importAllFrom(
		testResourceFilename("adium_test_data/Accounts.plist"),
		testResourceFilename("adium_test_data/libpurple/accounts.xml"),
		testResourceFilename("adium_test_data/libpurple/prefs.xml"),
		testResourceFilename("adium_test_data/libpurple/blist.xml"),
		testResourceFilename("adium_test_data/otr.private_key"),
		testResourceFilename("adium_test_data/otr.fingerprints"),
	)

	c.Assert(ok, Equals, true)
	c.Assert(res, NotNil)
	c.Assert(len(res.Accounts), Equals, 3)

	c.Assert(*res.Accounts[0], DeepEquals, config.Account{
		Account:                 "baarfoo@coyim.com",
		Server:                  "talk.google.com",
		Port:                    5222,
		PrivateKeys:             [][]byte{[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x87, 0xc1, 0x7f, 0x67, 0x46, 0x72, 0x28, 0xe1, 0x78, 0x33, 0x5a, 0xa7, 0xe8, 0x9e, 0x68, 0x75, 0xac, 0xc, 0x29, 0x2d, 0xb, 0x6b, 0xfe, 0x40, 0xc9, 0xbd, 0x16, 0xe3, 0x5d, 0x0, 0x11, 0x5, 0xd2, 0xd0, 0x59, 0x48, 0xce, 0x99, 0xcf, 0xaa, 0x48, 0xf6, 0xda, 0x7e, 0xe3, 0xaf, 0x82, 0x8d, 0xc3, 0x1f, 0xf8, 0xc2, 0xfb, 0x93, 0xbf, 0xb7, 0xb2, 0xfe, 0x76, 0x86, 0x2, 0x39, 0xfd, 0x2, 0x74, 0xf7, 0xdb, 0x95, 0x46, 0x0, 0x8f, 0xde, 0x2e, 0x97, 0xeb, 0x90, 0x2a, 0xe, 0xad, 0x9, 0x97, 0x2d, 0x3b, 0x1a, 0x5, 0x37, 0xb1, 0x43, 0x80, 0xa4, 0x74, 0x46, 0x7, 0x83, 0xa, 0x99, 0xf1, 0x0, 0xcc, 0x36, 0xb, 0xd8, 0x16, 0x9c, 0xce, 0x9c, 0x19, 0x62, 0x7a, 0x31, 0x27, 0xec, 0xbc, 0xf, 0xdb, 0x50, 0xd4, 0xf8, 0xe9, 0x75, 0x50, 0x69, 0xe2, 0xb, 0x82, 0x82, 0x3, 0x3, 0x0, 0x0, 0x0, 0x14, 0x8d, 0x32, 0xdb, 0xfd, 0x90, 0xde, 0x65, 0x7a, 0xaf, 0xd1, 0x4f, 0xfc, 0xd3, 0xb2, 0x1a, 0x7f, 0xa3, 0x98, 0x45, 0x49, 0x0, 0x0, 0x0, 0x80, 0x13, 0x1c, 0xd5, 0xa2, 0xe1, 0x9c, 0x1e, 0xea, 0x82, 0xb5, 0xad, 0x6e, 0x5d, 0x9c, 0x63, 0x52, 0x58, 0x17, 0xc3, 0xb3, 0x99, 0x50, 0xac, 0x1f, 0x4b, 0x4a, 0x1c, 0x1e, 0xee, 0xd0, 0x9a, 0xe9, 0x5d, 0x6, 0xf6, 0x3a, 0x57, 0x19, 0x95, 0xf9, 0xb9, 0xff, 0x4e, 0x7, 0xe3, 0xfc, 0xdd, 0xc0, 0xfc, 0x97, 0xee, 0x88, 0xa5, 0xf6, 0x48, 0xa9, 0x30, 0x80, 0x5e, 0xf7, 0x34, 0xf4, 0xed, 0x29, 0xe7, 0x18, 0xaf, 0x93, 0x9a, 0x76, 0x6b, 0xc5, 0x4b, 0x5f, 0x9b, 0x43, 0xce, 0x3e, 0x70, 0x33, 0x99, 0xd7, 0xb1, 0xa6, 0x8e, 0x4b, 0x7c, 0xb0, 0x23, 0x9a, 0x42, 0xee, 0x2c, 0x68, 0xb0, 0x6f, 0xe2, 0xb5, 0xab, 0x59, 0xf7, 0xa9, 0x26, 0xaf, 0x96, 0xed, 0xaa, 0xe6, 0x86, 0x95, 0x43, 0x78, 0x63, 0xe7, 0x6e, 0xa6, 0x90, 0x39, 0xcd, 0x76, 0x92, 0xa, 0x83, 0x7b, 0xc4, 0x6f, 0x1b, 0x38, 0x0, 0x0, 0x0, 0x80, 0x3, 0x85, 0xe6, 0xcc, 0x5, 0xe5, 0x1b, 0x4a, 0x3f, 0x45, 0xdd, 0xc8, 0x58, 0xec, 0x4c, 0x77, 0x9, 0x99, 0x47, 0xf1, 0x88, 0x8b, 0x6e, 0xe4, 0x26, 0xf3, 0xc4, 0x35, 0x69, 0xbd, 0xf2, 0xc, 0xbb, 0xa6, 0xe5, 0x50, 0x6, 0xec, 0xdd, 0x98, 0xf4, 0x53, 0xfa, 0x20, 0xf0, 0x6c, 0x38, 0xe3, 0xf3, 0x39, 0x6d, 0x8a, 0x3f, 0x40, 0xea, 0x50, 0xac, 0xd9, 0x59, 0x30, 0xa3, 0xb4, 0xf8, 0xf2, 0x79, 0xb8, 0x65, 0x56, 0xca, 0xab, 0x7e, 0xe6, 0xdc, 0x55, 0xe1, 0x76, 0x3e, 0x2f, 0x5d, 0x36, 0x50, 0xc, 0xf6, 0x72, 0xef, 0x94, 0x14, 0x96, 0x32, 0xdc, 0x49, 0x91, 0xc0, 0x25, 0x86, 0x88, 0x7a, 0x39, 0x67, 0x27, 0x46, 0x0, 0x40, 0x3c, 0xb8, 0x78, 0x90, 0xfc, 0x9, 0x69, 0x8a, 0x47, 0xf2, 0x50, 0xb2, 0x1d, 0xbe, 0x46, 0x62, 0x44, 0x58, 0x21, 0x6a, 0xe9, 0x5a, 0x6, 0x47, 0x11, 0x0, 0x0, 0x0, 0x14, 0x39, 0x41, 0x1a, 0x59, 0x74, 0x58, 0x75, 0xf, 0x7a, 0x5f, 0x2b, 0xd5, 0x61, 0x85, 0xdf, 0x71, 0x93, 0xf, 0xd4, 0x2e}},
		LegacyKnownFingerprints: nil,
		Peers: []*config.Peer{
			&config.Peer{
				UserID: "abcaaaa@mesopotamia.xmpp",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("4bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
				},
			},
			&config.Peer{
				UserID: "fooba@mesopotamia.xmpp",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("5bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     false,
					},
					&config.Fingerprint{
						Fingerprint: decode("6bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
				},
			},
			&config.Peer{
				UserID: "hmmabc@mesopotamia.xmpp",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("2bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
					&config.Fingerprint{
						Fingerprint: decode("3bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     false,
					},
				},
			},
		},
		HideStatusUpdates:    false,
		Proxies:              []string{},
		OTRAutoTearDown:      false,
		OTRAutoAppendTag:     false,
		OTRAutoStartSession:  false,
		AlwaysEncrypt:        false,
		InstanceTag:          0x0,
		ConnectAutomatically: false})

	c.Assert(*res.Accounts[1], DeepEquals, config.Account{
		Account:              "baarfoo@coyim.org",
		Proxies:              []string{},
		Port:                 5222,
		HideStatusUpdates:    false,
		OTRAutoTearDown:      false,
		OTRAutoAppendTag:     false,
		OTRAutoStartSession:  false,
		AlwaysEncrypt:        false,
		InstanceTag:          0x0,
		ConnectAutomatically: false})

	expectedAccount := config.Account{
		Account:                 "foobaar@mesopotamia.xmpp",
		Server:                  "talk.google.com",
		Proxies:                 []string{},
		Port:                    5222,
		PrivateKeys:             [][]byte{[]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x87, 0xc1, 0x7f, 0x67, 0x46, 0x72, 0x28, 0xe1, 0x78, 0x33, 0x5a, 0xa7, 0xe8, 0x9e, 0x68, 0x75, 0xac, 0xc, 0x29, 0x2d, 0xb, 0x6b, 0xfe, 0x40, 0xc9, 0xbd, 0x16, 0xe3, 0x5d, 0x0, 0x11, 0x5, 0xd2, 0xd0, 0x59, 0x48, 0xce, 0x99, 0xcf, 0xaa, 0x48, 0xf6, 0xda, 0x7e, 0xe3, 0xaf, 0x82, 0x8d, 0xc3, 0x1f, 0xf8, 0xc2, 0xfb, 0x93, 0xbf, 0xb7, 0xb2, 0xfe, 0x76, 0x86, 0x2, 0x39, 0xfd, 0x2, 0x74, 0xf7, 0xdb, 0x95, 0x46, 0x0, 0x8f, 0xde, 0x2e, 0x97, 0xeb, 0x90, 0x2a, 0xe, 0xad, 0x9, 0x97, 0x2d, 0x3b, 0x1a, 0x5, 0x37, 0xb1, 0x43, 0x80, 0xa4, 0x74, 0x46, 0x7, 0x83, 0xa, 0x99, 0xf1, 0x0, 0xcc, 0x36, 0xb, 0xd8, 0x16, 0x9c, 0xce, 0x9c, 0x19, 0x62, 0x7a, 0x31, 0x27, 0xec, 0xbc, 0xf, 0xdb, 0x50, 0xd4, 0xf8, 0xe9, 0x75, 0x50, 0x69, 0xe2, 0xb, 0x82, 0x82, 0x3, 0x3, 0x0, 0x0, 0x0, 0x14, 0x8d, 0x32, 0xdb, 0xfd, 0x90, 0xde, 0x65, 0x7a, 0xaf, 0xd1, 0x4f, 0xfc, 0xd3, 0xb2, 0x1a, 0x7f, 0xa3, 0x98, 0x45, 0x49, 0x0, 0x0, 0x0, 0x80, 0x13, 0x1c, 0xd5, 0xa2, 0xe1, 0x9c, 0x1e, 0xea, 0x82, 0xb5, 0xad, 0x6e, 0x5d, 0x9c, 0x63, 0x52, 0x58, 0x17, 0xc3, 0xb3, 0x99, 0x50, 0xac, 0x1f, 0x4b, 0x4a, 0x1c, 0x1e, 0xee, 0xd0, 0x9a, 0xe9, 0x5d, 0x6, 0xf6, 0x3a, 0x57, 0x19, 0x95, 0xf9, 0xb9, 0xff, 0x4e, 0x7, 0xe3, 0xfc, 0xdd, 0xc0, 0xfc, 0x97, 0xee, 0x88, 0xa5, 0xf6, 0x48, 0xa9, 0x30, 0x80, 0x5e, 0xf7, 0x34, 0xf4, 0xed, 0x29, 0xe7, 0x18, 0xaf, 0x93, 0x9a, 0x76, 0x6b, 0xc5, 0x4b, 0x5f, 0x9b, 0x43, 0xce, 0x3e, 0x70, 0x33, 0x99, 0xd7, 0xb1, 0xa6, 0x8e, 0x4b, 0x7c, 0xb0, 0x23, 0x9a, 0x42, 0xee, 0x2c, 0x68, 0xb0, 0x6f, 0xe2, 0xb5, 0xab, 0x59, 0xf7, 0xa9, 0x26, 0xaf, 0x96, 0xed, 0xaa, 0xe6, 0x86, 0x95, 0x43, 0x78, 0x63, 0xe7, 0x6e, 0xa6, 0x90, 0x39, 0xcd, 0x76, 0x92, 0xa, 0x83, 0x7b, 0xc4, 0x6f, 0x1b, 0x38, 0x0, 0x0, 0x0, 0x80, 0x3, 0x85, 0xe6, 0xcc, 0x5, 0xe5, 0x1b, 0x4a, 0x3f, 0x45, 0xdd, 0xc8, 0x58, 0xec, 0x4c, 0x77, 0x9, 0x99, 0x47, 0xf1, 0x88, 0x8b, 0x6e, 0xe4, 0x26, 0xf3, 0xc4, 0x35, 0x69, 0xbd, 0xf2, 0xc, 0xbb, 0xa6, 0xe5, 0x50, 0x6, 0xec, 0xdd, 0x98, 0xf4, 0x53, 0xfa, 0x20, 0xf0, 0x6c, 0x38, 0xe3, 0xf3, 0x39, 0x6d, 0x8a, 0x3f, 0x40, 0xea, 0x50, 0xac, 0xd9, 0x59, 0x30, 0xa3, 0xb4, 0xf8, 0xf2, 0x79, 0xb8, 0x65, 0x56, 0xca, 0xab, 0x7e, 0xe6, 0xdc, 0x55, 0xe1, 0x76, 0x3e, 0x2f, 0x5d, 0x36, 0x50, 0xc, 0xf6, 0x72, 0xef, 0x94, 0x14, 0x96, 0x32, 0xdc, 0x49, 0x91, 0xc0, 0x25, 0x86, 0x88, 0x7a, 0x39, 0x67, 0x27, 0x46, 0x0, 0x40, 0x3c, 0xb8, 0x78, 0x90, 0xfc, 0x9, 0x69, 0x8a, 0x47, 0xf2, 0x50, 0xb2, 0x1d, 0xbe, 0x46, 0x62, 0x44, 0x58, 0x21, 0x6a, 0xe9, 0x5a, 0x6, 0x47, 0x11, 0x0, 0x0, 0x0, 0x14, 0x29, 0x41, 0x1a, 0x59, 0x74, 0x58, 0x75, 0xf, 0x7a, 0x5f, 0x2b, 0xd5, 0x61, 0x85, 0xdf, 0x71, 0x93, 0xf, 0xd4, 0x2e}},
		LegacyKnownFingerprints: nil,
		Peers: []*config.Peer{
			&config.Peer{
				UserID: "abc@jabber.org",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("8bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
				},
			},
			&config.Peer{
				UserID: "abdfdergsfdfgd@gmail.com",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("1bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
				},
			},
			&config.Peer{
				UserID: "llaall@gmail.com",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("7bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
				},
			},
			&config.Peer{
				UserID: "xxqqx@coy.im",
				Fingerprints: []*config.Fingerprint{
					&config.Fingerprint{
						Fingerprint: decode("10cdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
					&config.Fingerprint{
						Fingerprint: decode("9bcdefabcdefabcdefabcdefabcdefabcdefcdef"),
						Trusted:     true,
					},
				},
			},
		},
		HideStatusUpdates:    false,
		OTRAutoTearDown:      false,
		OTRAutoAppendTag:     false,
		OTRAutoStartSession:  false,
		AlwaysEncrypt:        false,
		InstanceTag:          0x0,
		ConnectAutomatically: false}

	c.Assert(*res.Accounts[2], DeepEquals, expectedAccount)
}
