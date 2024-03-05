module libcore

go 1.21.4

toolchain go1.22.0

require (
	github.com/pion/stun/v2 v2.0.0
	github.com/sirupsen/logrus v1.9.3
	github.com/ulikunitz/xz v0.5.11
	github.com/v2fly/v2ray-core/v5 v5.14.1
	golang.org/x/mobile v0.0.0-20240213143359-d1f7d3436075
	golang.org/x/net v0.22.0
	golang.org/x/sys v0.18.0
	google.golang.org/protobuf v1.32.0
	gvisor.dev/gvisor v0.0.0-20240212194444-1796ac7f1d01
)

require (
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/aead/cmac v0.0.0-20160719120800-7af84192f0b1 // indirect
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/boljen/go-bitmap v0.0.0-20151001105940-23cd2fb0ce7d // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/dgryski/go-metro v0.0.0-20200812162917-85c65e2d0165 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/pprof v0.0.0-20210407192527-94a9f03dee38 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/klauspost/reedsolomon v1.9.3 // indirect
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 // indirect
	github.com/mustafaturan/bus v1.0.2 // indirect
	github.com/mustafaturan/monoton v1.0.0 // indirect
	github.com/onsi/ginkgo/v2 v2.9.5 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/pion/dtls/v2 v2.2.8 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/sctp v1.7.6 // indirect
	github.com/pion/transport/v2 v2.2.4 // indirect
	github.com/pion/transport/v3 v3.0.1 // indirect
	github.com/pires/go-proxyproto v0.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/quic-go v0.41.0 // indirect
	github.com/refraction-networking/utls v1.6.3 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/sagernet/sing v0.3.4 // indirect
	github.com/sagernet/sing-shadowsocks v0.2.6 // indirect
	github.com/sagernet/sing-shadowsocks2 v0.2.0 // indirect
	github.com/secure-io/siv-go v0.0.0-20180922214919-5ff40651e2c4 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20220411075957-e3b120b3f5fb // indirect
	github.com/v2fly/BrowserBridge v0.0.0-20210430233438-0570fc1d7d08 // indirect
	github.com/v2fly/ss-bloomring v0.0.0-20210312155135-28617310f63e // indirect
	github.com/xiaokangwang/VLite v0.0.0-20231225174116-75fa4b06e9f2 // indirect
	github.com/xtaci/smux v1.5.15 // indirect
	github.com/xtls/reality v0.0.0-20231112171332-de1173cf2b19 // indirect
	go.uber.org/mock v0.4.0 // indirect
	go4.org/netipx v0.0.0-20231129151722-fdeea329fbba // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20230725093048-515e97ebf090 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.18.0 // indirect
	golang.zx2c4.com/wintun v0.0.0-20230126152724-0fa3db229ce2 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20231211153847-12269c276173 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.2.1 // indirect
)

replace github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 => github.com/xiaokangwang/struc v0.0.0-20231031203518-0e381172f248

//replace github.com/v2fly/v2ray-core/v5 => ../v2ray-core

replace github.com/v2fly/v2ray-core/v5 => github.com/dyhkwong/v2ray-core/v5 v5.14.2-0.20240305054242-059249efa11b
