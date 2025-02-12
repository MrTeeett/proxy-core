package sniff_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"

	"proxy-core/adapter"
	"proxy-core/common/sniff"
	C "proxy-core/constant"

	"github.com/stretchr/testify/require"
)

func TestSniffBittorrent(t *testing.T) {
	t.Parallel()

	packets := []string{
		"13426974546f7272656e742070726f746f636f6c0000000000100000e21ea9569b69bab33c97851d0298bdfa89bc90922d5554313631302dea812fcd6a3563e3be40c1d1",
		"13426974546f7272656e742070726f746f636f6c00000000001000052aa4f5a7e209e54b32803d43670971c4c8caaa052d5452333030302d653369733079647675763638",
		"13426974546f7272656e742070726f746f636f6c00000000001000052aa4f5a7e209e54b32803d43670971c4c8caaa052d5452343035302d6f7a316c6e79377931716130",
	}

	for _, pkt := range packets {
		pkt, err := hex.DecodeString(pkt)
		require.NoError(t, err)
		var metadata adapter.InboundContext
		err = sniff.BitTorrent(context.TODO(), &metadata, bytes.NewReader(pkt))
		require.NoError(t, err)
		require.Equal(t, C.ProtocolBitTorrent, metadata.Protocol)
	}
}

func TestSniffUTP(t *testing.T) {
	t.Parallel()

	packets := []string{
		"010041a282d7ee7b583afb160004000006d8318da776968f92d666f7963f32dae23ba0d2c810d8b8209cc4939f54fde9eeaa521c2c20c9ba7f43f4fb0375f28de06643b5e3ca4685ab7ac76adca99783be72ef05ed59ef4234f5712b75b4c7c0d7bee8fe2ca20ad626ba5bb0ffcc16bf06790896f888048cf72716419a07db1a3dca4550fbcea75b53e97235168a221cf3e553dfbb723961bd719fab038d86e0ecb74747f5a2cd669de1c4b9ad375f3a492d09d98cdfad745435625401315bbba98d35d32086299801377b93495a63a9efddb8d05f5b37a5c5b1c0a25e917f12007bb5e05013ada8aff544fab8cadf61d80ddb0b60f12741e44515a109d144fd53ef845acb4b5ccf0d6fc302d7003d76df3fc3423bb0237301c9e88f900c2d392a8e0fdb36d143cf7527a93fd0a2638b746e72f6699fffcd4fd15348fce780d4caa04382fd9faf1ca0ae377ca805da7536662b84f5ee18dd3ae38fcb095a7543e55f9069ae92c8cf54ae44e97b558d35e2545c66601ed2149cbc32bd6df199a2be7cf0da8b2ff137e0d23e776bc87248425013876d3a3cc31a83b424b752bd0346437f24b532978005d8f5b1b0be1a37a2489c32a18a9ad3118e3f9d30eb299bffae18e1f0677c2a5c185e62519093fe6bc2b7339299ea50a587989f726ca6443a75dd5bb936f6367c6355d80fae53ff529d740b2e5576e3eefdf1fdbfc69c3c8d8ac750512635de63e054bee1d3b689bc1b2bc3d2601e42a00b5c89066d173d4ae7ffedfd2274e5cf6d868fbe640aedb69b8246142f00b32d459974287537ddd5373460dcbc92f5cfdd7a3ed6020822ae922d947893752ca1983d0d32977374c384ac8f5ab566859019b7351526b9f13e932037a55bb052d9deb3b3c23317e0784fdc51a64f2159bfea3b069cf5caf02ee2c3c1a6b6b427bb16165713e8802d95b5c8ed77953690e994bd38c9ae113fedaf6ee7fc2b96c032ceafc2a530ad0422e84546b9c6ad8ef6ea02fa508abddd1805c38a7b42e9b7c971b1b636865ebec06ed754bb404cd6b4e6cc8cb77bd4a0c43410d5cd5ef8fe853a66d49b3b9e06cb141236cdbfdd5761601dc54d1250b86c660e0f898fe62526fdd9acf0eab60a3bbbb2151970461f28f10b31689594bea646c4b03ee197d63bdef4e5a7c22716b3bb9494a83b78ecd81b338b80ac6c09c43485b1b09ba41c74343832c78f0520c1d659ac9eb1502094141e82fb9e5e620970ebc0655514c43c294a7714cbf9a499d277daf089f556398a01589a77494bec8bfb60a108f3813b55368672b88c1af40f6b3c8b513f7c70c3e0efce85228b8b9ec67ba0393f9f7305024d8e2da6a26cf85613d14f249170ce1000089df4c9c260df7f8292aa2ecb5d5bac97656d59aa248caedea2d198e51ce87baece338716d114b458de02d65c9ff808ca5b5b73723b4d1e962d9ac2d98176544dc9984cf8554d07820ef3dd0861cfe57b478328046380de589adad94ee44743ffac73bb7361feca5d56f07cf8ce75080e261282ae30350d7882679b15cab9e7e53ddf93310b33f7390ae5d318bb53f387e6af5d0ef4f947fc9cb8e7e38b52c7f8d772ece6156b38d88796ea19df02c53723b44df7c76315a0de9462f27287e682d2b4cda1a68fe00d7e48c51ee981be44e1ca940fb5190c12655edb4a83c3a4f33e48a015692df4f0b3d61656e362aca657b5ae8c12db5a0db3db1e45135ee918b66918f40e53c4f83e9da0cddfe63f736ae751ab3837a30ae3220d8e8e311487093a7b90c7e7e40dd54ca750e19452f9193aa892aa6a6229ab493dadae988b1724f7898ee69c36d3eb7364c4adbeca811cfe2065873e78c2b6dfdf1595f7a7831c07e03cda82e4f86f76438dfb2b07c13638ce7b509cfa71b88b5102b39a203b423202088e1c2103319cb32c13c1e546ff8612fa194c95a7808ab767c265a1bd5fa0efed5c8ec1701876a00ec8",
		"01001ecb68176f215d04326300100000dbcf30292d14b54e9ee2d115ee5b8ebc7fad3e882d4fcdd0c14c6b917c11cb4c6a9f410b52a33ae97c2ac77c7a2b122b8955e09af3c5c595f1b2e79ca57cfe44c44e069610773b9bc9ba223d7f6b383e3adddd03fb88a8476028e30979c2ef321ffc97c5c132bcf9ac5b410bbb5ec6cefca3c7209202a14c5ae922b6b157b0a80249d13ffe5b996af0bc8e54ba576d148372494303e7ead0602b05b9c8fc97d48508a028a04d63a1fd28b0edfcd5c51715f63188b53eefede98a76912dca98518551a8856567307a56a702cbfcc115ea0c755b418bc2c7b57721239b82f09fb24328a4b0ce0f109bcb2a64e04b8aadb1f8487585425acdf8fc4ec8ea93cfcec5ac098bb29d42ddef6e46b03f34a5de28316726699b7cb5195c33e5c48abe87d591d63f9991c84c30819d186d6e0e95fd83c8dff07aa669c4430989bcaccfeacb9bcadbdb4d8f1964dbeb9687745656edd30b21c66cc0a1d742a78717d134a19a7f02d285a4973b1a198c00cfdff4676608dc4f3e817e3463c3b4e2c80d3e8d4fbac541a58a2fb7ad6939f607f8144eff6c8b0adc28ee5609ea158987519892fb",
		"21001ecb6817f2805d044fd700100000dbd03029",
		"410277ef0b1fb1f60000000000040000c233000000080000000000000000",
	}

	for _, pkt := range packets {
		pkt, err := hex.DecodeString(pkt)
		require.NoError(t, err)
		var metadata adapter.InboundContext
		err = sniff.UTP(context.TODO(), &metadata, pkt)
		require.NoError(t, err)
		require.Equal(t, C.ProtocolBitTorrent, metadata.Protocol)
	}
}

func TestSniffUDPTracker(t *testing.T) {
	t.Parallel()

	connectPackets := []string{
		"00000417271019800000000078e90560",
		"00000417271019800000000022c5d64d",
		"000004172710198000000000b3863541",
	}

	for _, pkt := range connectPackets {
		pkt, err := hex.DecodeString(pkt)
		require.NoError(t, err)

		var metadata adapter.InboundContext
		err = sniff.UDPTracker(context.TODO(), &metadata, pkt)
		require.NoError(t, err)
		require.Equal(t, C.ProtocolBitTorrent, metadata.Protocol)
	}
}
