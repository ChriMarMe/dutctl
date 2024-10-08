# Changelog

## 0.1.0 (2024-09-26)


### Features

* add generated code for RPC service ([f3770a0](https://github.com/BlindspotSoftware/dutctl/commit/f3770a0199cbbfce477192c33995d659f0e9d562))
* add pkg/dut ([51027cc](https://github.com/BlindspotSoftware/dutctl/commit/51027ccb5d5471462afead94349a259a0cef72b9))
* add pkg/module with Module interface definition ([594224c](https://github.com/BlindspotSoftware/dutctl/commit/594224ce621fb2145f523975ccd47f29df3143ef))
* cli for dutagent ([1388faf](https://github.com/BlindspotSoftware/dutctl/commit/1388faf88d96cff39f3bb90caec0d9c505675b69))
* cli for dutctl ([80fc4bc](https://github.com/BlindspotSoftware/dutctl/commit/80fc4bc4022ab99a3338b5e0d8dde1bbfe5c8ceb))
* communication design: add 'Details'-RPC ([5c41ffa](https://github.com/BlindspotSoftware/dutctl/commit/5c41ffa32a927ef107d9f2080a38d4f1f2c0a879))
* draft command for dutctl ([8eed5d1](https://github.com/BlindspotSoftware/dutctl/commit/8eed5d176baf860abddc08b135c62cab20c4b545))
* draft commands for dutctl and dutagend ([5c80434](https://github.com/BlindspotSoftware/dutctl/commit/5c80434483ed619902b594793594af7adc6d219f))
* dutagent runs module de-init on clean-up ([64ecf79](https://github.com/BlindspotSoftware/dutctl/commit/64ecf79af24740c0e87ea04c12c86b6663786192))
* dutagent runs module init on startup ([c34757f](https://github.com/BlindspotSoftware/dutctl/commit/c34757f417935d145c5b4316c1d4cbefa8e87af8))
* dutctl: validate transmitted file names against cmdline content ([842ca42](https://github.com/BlindspotSoftware/dutctl/commit/842ca427dc7acf171e9f9e4c57da4561653e7044))
* file transfer between client and agent ([9b09125](https://github.com/BlindspotSoftware/dutctl/commit/9b0912557c09bff9e6aa9c63423aca0c746ec34c))
* multi-module support ([f3b5bb3](https://github.com/BlindspotSoftware/dutctl/commit/f3b5bb31679fcea294fc7b062c5b0c498b16987e))
* parse devices from YAML config ([62bad03](https://github.com/BlindspotSoftware/dutctl/commit/62bad0364c7f864aaf64bba51f902285dfc89b24))
* plug-in system for modules ([07f6338](https://github.com/BlindspotSoftware/dutctl/commit/07f6338c8daf5fac55fb9fe2ebcb9cfa94cfc663))
* protobuf spec ([21a406a](https://github.com/BlindspotSoftware/dutctl/commit/21a406ae0cfd0598a848d167e3b0ceadb36b8cdc))


### Bug Fixes

* chanio constructors return errors ([2737ff9](https://github.com/BlindspotSoftware/dutctl/commit/2737ff93f36ca1ea4a15f9ee8c40465eb3e72f88))
* dutagent example config: remove empty command ([d2a1f9c](https://github.com/BlindspotSoftware/dutctl/commit/d2a1f9ca20d5af8ce4a8925f8e8a8e4c740e89d3))
* dutagent module broker errors on bad file transfers ([8ca9292](https://github.com/BlindspotSoftware/dutctl/commit/8ca92925e0b9617cf0e95186cc87fcf5ac3730dd))
* synch terminiation of module execution ([83d6f65](https://github.com/BlindspotSoftware/dutctl/commit/83d6f6529f0ee673727581f0ed4d14cdbd5450ec))


### Documentation

* add flow charts for RPCs ([d624218](https://github.com/BlindspotSoftware/dutctl/commit/d6242183204273d06e150bc7fe55064f4f3bbf14))
* add golangci-lint to CONTRIBUTING.md ([2ffe20b](https://github.com/BlindspotSoftware/dutctl/commit/2ffe20bae7aafec751a12058c12e04cebe02f770))
* add NLNet logo to README.md ([9d0ce00](https://github.com/BlindspotSoftware/dutctl/commit/9d0ce006a4461fad434c0d3fb4861c16ed01e447))
* communication design ([33f2611](https://github.com/BlindspotSoftware/dutctl/commit/33f26114bd2edb0e641cf2f34825667a4f18e60d))
* finalize initial project documentation ([e90b05d](https://github.com/BlindspotSoftware/dutctl/commit/e90b05d7245ddb98ad41cbab2fac3f0d0d6c3c93))
* finalize module documentation ([d182e1f](https://github.com/BlindspotSoftware/dutctl/commit/d182e1fa84a44368cfc79e607a6937a7db0e3bfa))
* fine tune module guide ([db2c43e](https://github.com/BlindspotSoftware/dutctl/commit/db2c43ef97c4e006a5f761892c9c89f3400f662c))
* fix code block style ([f171a81](https://github.com/BlindspotSoftware/dutctl/commit/f171a81d782ff1e15a23f3428145d60aac05707a))
* fix reference in protobuf/README.md ([6cecfb2](https://github.com/BlindspotSoftware/dutctl/commit/6cecfb204258225898297d4695e7d22c3d53b14d))
* installation instructions in protobuf/README.md ([f338748](https://github.com/BlindspotSoftware/dutctl/commit/f338748f1ef0ae25c56da072a48e976e607aed56))
* update Contributing.md regarding conventional commits ([6d1f593](https://github.com/BlindspotSoftware/dutctl/commit/6d1f59345cd1927a0f7e56edbcd300e776b4ad13))
* update module docu ([c464e7b](https://github.com/BlindspotSoftware/dutctl/commit/c464e7b8b30d3a893c0c4977e0099d4589d9961a))
* update README ([c64effa](https://github.com/BlindspotSoftware/dutctl/commit/c64effa173aaa9a634b186f8f8a938b407504c88))


### Other Work

* adapt naming in dutctl.proto ([bad6d7e](https://github.com/BlindspotSoftware/dutctl/commit/bad6d7ec94ff3b023856fd174287b6563d314a52))
* add device list type to pkg/dut ([c2a6f40](https://github.com/BlindspotSoftware/dutctl/commit/c2a6f4021822de9e7b0424c68c4b5a997edab433))
* code style and linting ([69f623f](https://github.com/BlindspotSoftware/dutctl/commit/69f623f108d75ae231a56ca01a96090a34d10049))
* dutagent cli ([69c7d4e](https://github.com/BlindspotSoftware/dutctl/commit/69c7d4e129c1f78c09151cde9c6d4318f6811cc7))
* dutagent-to-module communication ([6b41105](https://github.com/BlindspotSoftware/dutctl/commit/6b4110525cb802d697fb60a1efafa49aa73067b8))
* dutagent's Run RPC handler being a state machine ([97e5343](https://github.com/BlindspotSoftware/dutctl/commit/97e5343206a359dbbd79970a45f74cfecd7bf680))


### Chore

* release 0.1.0 ([96f8ad7](https://github.com/BlindspotSoftware/dutctl/commit/96f8ad7e92b252f4ba5a567ea0fc7bad0ce3e7a4))