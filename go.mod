module libcore

go 1.21.0

require (
	github.com/Dreamacro/clash v1.18.0
	github.com/golang/protobuf v1.5.3
	github.com/pion/stun/v2 v2.0.0
	github.com/sirupsen/logrus v1.9.3
	github.com/ulikunitz/xz v0.5.11
	github.com/v2fly/v2ray-core/v5 v5.9.0
	golang.org/x/mobile v0.0.0-20231006135142-2b44d11868fe
	golang.org/x/net v0.17.0
	golang.org/x/sys v0.13.0
	gvisor.dev/gvisor v0.0.0-20231020174304-b8a429915ff1
)

//replace github.com/v2fly/v2ray-core/v5 => ../v2ray-core

replace github.com/v2fly/v2ray-core/v5 => github.com/dyhkwong/v2ray-core/v5 v5.9.1-0.20231025114612-56eabd36753d

require (
	github.com/Dreamacro/protobytes v0.0.0-20230617041236-6500a9f4f158 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/aead/cmac v0.0.0-20160719120800-7af84192f0b1 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/boljen/go-bitmap v0.0.0-20151001105940-23cd2fb0ce7d // indirect
	github.com/bufbuild/protocompile v0.6.0 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/dgryski/go-camellia v0.0.0-20191119043421-69a8a13fb23d // indirect
	github.com/dgryski/go-idea v0.0.0-20170306091226-d2fb45a411fb // indirect
	github.com/dgryski/go-metro v0.0.0-20211217172704-adc40b04c140 // indirect
	github.com/dgryski/go-rc2 v0.0.0-20150621095337-8a9021637152 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gaukas/godicttls v0.0.4 // indirect
	github.com/geeksbaek/seed v0.0.0-20180909040025-2a7f5fb92e22 // indirect
	github.com/go-chi/chi/v5 v5.0.10 // indirect
	github.com/go-chi/render v1.0.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.5 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jhump/protoreflect v1.15.3 // indirect
	github.com/kierdavis/cfb8 v0.0.0-20180105024805-3a17c36ee2f8 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.0.12 // indirect
	github.com/klauspost/reedsolomon v1.9.3 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 // indirect
	github.com/mustafaturan/bus v1.0.2 // indirect
	github.com/mustafaturan/monoton v1.0.0 // indirect
	github.com/onsi/ginkgo/v2 v2.10.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pion/dtls/v2 v2.2.7 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/sctp v1.7.6 // indirect
	github.com/pion/transport/v2 v2.2.1 // indirect
	github.com/pion/transport/v3 v3.0.1 // indirect
	github.com/pires/go-proxyproto v0.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-20 v0.3.4 // indirect
	github.com/quic-go/quic-go v0.39.1 // indirect
	github.com/refraction-networking/utls v1.5.4 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sagernet/sing v0.2.14 // indirect
	github.com/sagernet/sing-shadowsocks v0.2.5 // indirect
	github.com/sagernet/sing-shadowsocks2 v0.1.4 // indirect
	github.com/secure-io/siv-go v0.0.0-20180922214919-5ff40651e2c4 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20220411075957-e3b120b3f5fb // indirect
	github.com/v2fly/BrowserBridge v0.0.0-20210430233438-0570fc1d7d08 // indirect
	github.com/v2fly/ss-bloomring v0.0.0-20210312155135-28617310f63e // indirect
	github.com/xiaokangwang/VLite v0.0.0-20220418190619-cff95160a432 // indirect
	github.com/xtaci/smux v1.5.15 // indirect
	github.com/xtls/reality v0.0.0-20231010003020-da6c695a34db // indirect
	go.starlark.net v0.0.0-20231016134836-22325403fcb3 // indirect
	go.uber.org/mock v0.3.0 // indirect
	go4.org/netipx v0.0.0-20230824141953-6213f710f925 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20230725093048-515e97ebf090 // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.2.1 // indirect
)
