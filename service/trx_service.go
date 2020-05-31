package service

import "github.com/vincecfl/go-common/log"

func TransferTrxHandle() {
	list := make([]*Account, 0)

	// 0
	//list = append(list, &Account{"TPMXwJuAt6Nbi5M3njhhUsEHw5QUCx5VV2", "1b93c7501b89d1e239fcfac50270036863271b31736c7960984a6be2c02c902a"})
	//list = append(list, &Account{"TLcLcDTDdsKc1ZRDFgt6JPKYyquX9v25V4", "4fcf0b8df38091c9908ee52f75edaa9bf16d55c9ade992dd043a8fb684efa663"})

	// 1
	//list = append(list, &Account{"TKeHVce6sxDdG7SbNfN81v1PD9nbTGeSXH", "27c8c9144a576820e35d2ecd3b277a9b264430d90903301376c3aa914db4ca14"})
	//list = append(list, &Account{"TPfaaP5WYt5L8TLtkVYzQXVvLeVPHqR8cX", "e873daf5b5549493bb38299e404e6d3442d814e5aad6d4d27d043190e02ca814"})
	//list = append(list, &Account{"TTuwCHR9ofuV6k2wheBLf4wDEaMcyM5Dk7", "785473388246e048645488f52eb5fcaa2b80f0a23e3f33920b04e23cb98351d5"})

	// 2
	//list = append(list, &Account{"TBFpBBrAgKvSm1cRK9EfwPXCupSbqH3yXB", "A95D7AC43F0360FCDE2F77380D492683A38E474B87E71088503DFE1E6F4473FB"})
	//list = append(list, &Account{"TXULv92s3eMMpeyFoUveVsAwqbKqMrzW1x", "BA04FB6150F4E2A3FE6FE383D1A782E6540B4AA9C89ABB9E68DE536B7E68CE21"})
	//list = append(list, &Account{"TDkyCoEo2XG3HFa21b4YGC4szqBaHoDHjj", "2294BA50616CDF6065E055BE81F06C12651BD53F71E57D7164F11E05AE1EDAAB"})

	// 3
	//list = append(list, &Account{"TTCT6rgLXGnoJwBuMrVqXjRyyAFgdtgHro", "A6662FE759D9FF76143D5635D9D607D148A3604A5C64C4E34454D6BE9215E3C5"})
	//list = append(list, &Account{"TJmmJEfCGvZMYe9FSZ2niVFX1omx43HnzK", "A45876643F9C7F573F29E7FA34294637A1DCA210B4565D69219A9EBAF11C5549"})
	//list = append(list, &Account{"TUkWgBw178XuimVCkSPqp2Mjyrfpp9ApKt", "5B6AAB5711FC4C27F8C970708CE38A823D76EF9996E8BBDE9547B17354871849"})

	// 4
	//list = append(list, &Account{"TPwqzXcUgQaFTFMU6AefcbKedB2u1D1qiV", "427CAFECEB52553FBF3B31119633E26408AE884323C3EDB13A71FFC82E4346B0"})
	//list = append(list, &Account{"TSxtvS7DTmZGvm3nnxmu6DysCZ8a1JtG7E", "53B3AB7ADAB16110A37450F6F19DEA893C5A2D31DDED098488E95CF0ACB27EEE"})
	//list = append(list, &Account{"TXGSMzBT9RuW12t3rdP7tc615UBmGRWBXS", "3BF6E7522C01515D688FC117BBD20371A6C4D2A8A4651D70D69D6505C2D48F94"})

	// 5
	//list = append(list, &Account{"TDo2Sg79RBuri8j839M5a3N3M8VkVqPT64", "5CCF93D74CE566B3857267C4A92CBDE46A5C68E84A0BC468DC345825345C85A2"})
	//list = append(list, &Account{"TRdTi9rTKAdaWDKZRU9aeBV9ii91gQ3KKs", "DC66A127B9150CE1821253F0DB8B3F2E0901913BFF255B643C5AE079D1DA5B03"})
	//list = append(list, &Account{"TPvB6nRgcvTxvMxs3yxi7UQQjh2ex4UoYD", "508CC008EC2CB5B98C742FB30326D41AB2CA8A4A3EDCC4CE7EBEB3457E3529CA"})

	// 1
	list = append(list, &Account{"TJPsEEvmeDoDuHqEncDJHonWgtugqPZMnw", "05ac964b6c4ab529dde4911a5061a75b7af8aa329d22d17af38a1226752276b2"})
	list = append(list, &Account{"TNy9qMBKEQx42NRJFPvRojYWjQ6TZYsaWR", "a1588598cd502c64d4f4aaa44eae835146d7aea71e7f6f53ec5853d6ed7d77a5"})
	list = append(list, &Account{"TPWeaEZwAdYcXZWQvZQtrHPH7i8WYd8gyW", "13030bc7767cad09d07ca5b3026afa01c5e913e9d18a0afc8cf8713919d91f35"})

	// 2
	list = append(list, &Account{"TKv9wmaVx2AdSVfhQPUagk9YVUiXSR2vag", "5674392a12d229ee3e734018f001c2b34003b4f29cb41ef893eb0eccad3d2617"})
	list = append(list, &Account{"TMqdtFHvVDjbiAwWvntSJHmMk9Y5CqoJuB", "df70943c0d26ca821647b6449ae3a293a3f90cc2b38e86f71729420fecefc5cf"})
	list = append(list, &Account{"TLDSS4vd2wtMGAQohgP6147KYSnNpoRNRP", "2f65986770d2a45450549d5ca931b3559b7de2aea5b8a20b1a8d2b3237e5701e"})

	// 3
	list = append(list, &Account{"TCTaeJdEvETjEMjcCMXS9vPpWqB7SATKtB", "2e874eda217e54427d235383f810e355f51d0e04d7f776fc986f133d1e07270e"})
	list = append(list, &Account{"TGxapmiedvYuKAxxuSnnGwJLbJYEzHt95D", "742d84534e0e96d10d16bbb6b4dac9b8cd095a97679881242c481a3b0c19760d"})
	list = append(list, &Account{"TWdt48GQPX4kQ6fsayiYMKewNWgJw27oWh", "e8c7bec1273ab16d4c99773874aa7cfb06ca4a30b960cac2f07dc4d3d85af735"})

	// 4
	list = append(list, &Account{"TKRp7GwWixwXpDxoBcBN1RjutfXLEEHJVQ", "6e8c92223a25ccdecc697b28996babf4de6bb9f1da1b48bf669855b09cff1ed1"})
	list = append(list, &Account{"TPQ8Pm9v9x3TeNjAzw6T46RmuvePekitZW", "daa6bb58fd177ebb932dfcdc6bdd57a3b49494bf823cf3fb6d2569551aeb1988"})
	list = append(list, &Account{"TWQXoQaViGBBkt7MAezchWmENrA5TvpXDC", "ed38880022ede0b9116ad61468d5948523555142e39a91f61fb914496fe194f9"})

	// 5
	list = append(list, &Account{"TGiPvpUcVZ9wL6PUkAXwsPfgUF467QQyUS", "ed94f27b3d594cc1e1aafc5ff394d9aceb49de1949cf921494ddad2541cf3cc7"})
	list = append(list, &Account{"TDewqsGspBMDK4Ffovy65iyorPeyegqq8Z", "90316beec09ea1b766162ed83886f440ea3f80d9f99216f12bb5a03b3138daaa"})
	list = append(list, &Account{"TVJ5564EBDqEW5Ry83z4Z4xSob2wSzEmrA", "6ba5ce5a7c95f734698fc97ace83cb3ff3d8ebd0bc6f59d69d9a9330b18c1104"})

	// 6
	list = append(list, &Account{"TK47EP9EReRXnj3Fi7p1YgDiJcTkuFnFSo", "487ffe37c5f8f8563d0bb918a9efefd5e8ec99caa8892a7cd2190e17cc7b4ec1"})
	list = append(list, &Account{"TBzpUf4GaCD9rg9erh9AUxYN3arMYczLri", "397c46297e532d4b7bfece5c4d5f72b0bb58328ad479fb1d3ece07621defa416"})
	list = append(list, &Account{"TA5VnifqtMSdFNynEvtHp7YUTbsJhafQ3t", "df8b0c83b4bf27643066284bc3e3ee5467684d0cfee5bdb4fc1250cd53584a74"})

	// 7
	list = append(list, &Account{"TPNMByGpWA5hKMDeGtsS32wQ8gVf996pTd", "a9317d7f8ffbcfbf54a7771288f17df085126707444d9e71797d54ab99053041"})
	list = append(list, &Account{"TZCBAKwzhfFfyMXzbYBY7jmyV3Pf3ardQW", "43a84924b55988b0dda1e5614d667fe6a5187c4dd6f5eb811b1d2b7be7e08d78"})
	list = append(list, &Account{"TKzuhKMG7KUcvgBupvo4eVHJEKuYStWNL1", "86f86dd622f2ba33d72bca4cb2b94f046ebfb53343a8110016060bcbd68253a5"})

	// 8
	list = append(list, &Account{"TMrpWFtTYj9fJfBvrhr62xjgm58k1SSe2a", "ee35048c9d0d82df1ee385446b811ef7a021ea300926862bb9b108a965b06174"})
	list = append(list, &Account{"TCtYw2A4Ch4vEgfAMKKA1mKPeR5Bz8qBzS", "47f964afe8ff4606ea2f218145c6d5dbd54441a8101c3a11020ad70407f12875"})
	list = append(list, &Account{"TTVFzwoUNfYB8AWPwnXnesFVoKTLHMGG3r", "22b89c91d8a96b23909fad05372c92fa9328b97eee19c8e4f994dc27c90a3700"})

	// 9
	list = append(list, &Account{"TKMatxLDzM1WLo26LyBaBMwFEJHF1Gw48F", "92e4685d7ec1d507b28146b3e184c815b5900e7107c12b2f1155bb235de8461f"})
	list = append(list, &Account{"TH7xVB5zYR82sXZfRKEGDPUdkHbywYMnYv", "e0bd18638e9a12a24ce8137c18b5ab90778b6c417cfb1f3e23bdd16647b4f978"})
	list = append(list, &Account{"TSFT7oHjSpoaCThVUiXWjDGadL5ni4wC3w", "998fc63c022e4a14889d49a5864c8291ca7122504dc93c5877c87803a3d22768"})

	// 10
	list = append(list, &Account{"TNnBLRx2jicWZjdfFwKKAfYcqFgH1cA5Ys", "00b1c8e65d5152b8c21bc98a5946ec71a1f3759007e7a2d9f4cc724bd055e847"})
	list = append(list, &Account{"TLu9b5SadtjgAFbw6McnKLew4swUJuk3Do", "bd40e26e28053d93b633a22b4e225fee0bc779ea12ecaab4610cd013348f2614"})
	list = append(list, &Account{"TYRiWhiYVr3MRgNXoTAuZg5QHGjhCuUAVN", "a7103c570134d029d43952472ca530e53430b947fad98a057ec014328ef775fb"})

	// 11
	list = append(list, &Account{"TTheWb1NiwywVgNVRFuqxK6xAFDG5eJfFP", "2f950d1e91848ac75149b0371836f44144fccf8277d29541c8f9b0f6be161555"})
	list = append(list, &Account{"TGp61Vcjq5VpHMXJoZRFxVz4cm6fzvuGPh", "567df6044b2871028783c6adb9ab3eee3f8b77b3cce9e39d917798308e49ed34"})
	list = append(list, &Account{"TPFuxdcFheJUTvw2udSoDtTC6eNxgE8FBi", "be89af001e2e05f13c529e117628b9ee66ffaf3f8c0b83c462fc2dfedca7ec03"})

	// 12
	list = append(list, &Account{"TUoLRzgy3Hxo14o3irY9CfjGxAPGYazt55", "c6fbc7bdde56f130690bc46061bd978071d929e8abc64b11fc586f8d89338d4a"})
	list = append(list, &Account{"TEPjXt9RbzaLjemrf78v7SdVofWV5DJYkQ", "c87257e98c5735f24ebcc5aaef7fcc21895d03710f9b510202f94734b8147763"})
	list = append(list, &Account{"TJKEaPzZSiQ9W8KnSCK8JsUz3Bfo8Lq4dQ", "f8e3ec713e14b067860d7ab4f595c4a9ed0f14021a046fde42f524708169f751"})

	// 13
	list = append(list, &Account{"TKWKAFWigKaX9tQYqkFMBEfgp9UqcJqaLj", "77aae2b78095cedccfca415fe13c57d8dc7b61fff9379a230dac551c11422e49"})
	list = append(list, &Account{"TDGBCUfQPN6bZcYDA2dnH8qCHdgGn3zs72", "f8573e7f55cca050bc557721c97fcc8c389f06aa27e7ff4c38a444884a437060"})
	list = append(list, &Account{"TPDAbTUQ8bXRGvFZWbHyNPpRBJHYLhtrA4", "ee614d8d0049f1e7f6ba303b3c9b0023ca987861d1894779543a96188eadc35c"})

	// 14
	list = append(list, &Account{"TAdJJtb2BoxvYzkD3fKxpMeJKdD3sDPXEL", "687f6947e28054adff55f87f0a9f23ea40f25db97e40b499336da67b6392067f"})
	list = append(list, &Account{"TCqJb7BWujnJD4g9KCTtUB3Lnbdqdwb1R9", "4efc6e4061bb1c35422168224909cf383b2b86b997a9b85246d5cedaba95a673"})
	list = append(list, &Account{"TGCZawdxM5ten5RADDvtwY7pU2o74hWEqr", "93cdf9e2c1831402084dea3daf3703909b2cce0f62ac10aaf766a9f76755e39b"})

	// 15
	list = append(list, &Account{"TDpn55itxqhQAYCuUsVhgThycRxdLewWfd", "78f85e07d830a4d92637477f2110c87efb215c6cba7cc45b6891cefa08e76190"})
	list = append(list, &Account{"TMxRzFsanjG8igURYfwZWpQMAcYYJLbg9M", "e074eacb4e4ad70458e4a665251c6be634278a3762d93ff066d0803e66ff5614"})
	list = append(list, &Account{"TXjKhA1bLFP1spKFYgkPF7XokVfEjZWha7", "f10ba92d045c56f687a3c1396b4f950798e088eee74e021eda5651ff5a46a714"})

	// 16
	list = append(list, &Account{"TXbwfYSqKsKi6YjM5WsxVQyYVMNB91Luak", "4cf59c7ec3026be89f4e7a876b639f3e9110fb8a7c8430fac4d3ece2533ef716"})
	list = append(list, &Account{"TQX1R3tCM8aJLbHKZokJeB6xwXgSzJFCCS", "7bc079a58f5dbba07b0735d78f68a93e5fc37a22b7456da36325e9ccc8bbc8d6"})
	list = append(list, &Account{"TDwZhqqRJr9YZHPiDyshYN6ALxtRgDfNn2", "a35a83a7ec96375306a737e5825edf2caddc139e7da593224e1aff3e23cf2925"})

	// 17
	list = append(list, &Account{"TSdSUvGFTvYAnHv5FUabz4aYZCBBrNrCWk", "d5e6562a4adc8ba9f4c088d0b3052db01d35e05a022f737ab1697f6612f86ce7"})
	list = append(list, &Account{"TNars18p862CQx7xCnoHUDo7izeRbprsGh", "33b1f566c060465a1d8dc9c6ce94cf710e46bbbdc377249be2a85ad86a5e5291"})
	list = append(list, &Account{"TQokvUc1mPU5W5D2JMVa3RAjYcjxNoBbjk", "ddd95cbbc8c4ce34adb075f8dddaba1d2ad828a118a822db26c899a459cf226a"})

	// 18
	list = append(list, &Account{"TCN1X9rRf9X6bm9u6S2Vqn8Nm3ejmt2suP", "b86d76e399eed295be10e833525270f16b759021ccad1aedc7d83d17dcaca97d"})
	list = append(list, &Account{"TSK6fjterf2FYVWwXmPv4jCVNyRe2ioGd2", "c692a58be6cb9ce7975683e39079c5a82ccf3bf5cbee86d2181a99889b1aa4fc"})
	list = append(list, &Account{"TU9bZu9dxFd9rteKiMG3oEYWjQHjyHeccK", "1d92b332111fad5d99c7c99d6b8fbd726f1fbbeca4d110cbafc9349de063a524"})

	// 19
	list = append(list, &Account{"TKsTLrszLQL9WK2UHrFHG7XYn8rJdCJSMx", "04172ca0e352772c9e04c0ff939452b0a734db191d327646ccfa5f7a07073870"})
	list = append(list, &Account{"TWzSKckwdKXUNiEfoi5Yr6EkVZ5Rf8oUZu", "fc3f0a49b09017584c1445db0339b8448d1a3f06905975a7dfbe0f35462aac3e"})
	list = append(list, &Account{"TUSRzfgchYprht8XZGPEU1faN86z8P1xG6", "670b30c57684130964beafd40986c97b7dfb0801e55612e2120f5c266ab4ea84"})

	// 20
	list = append(list, &Account{"TFgesed4L5ZSFzFqHWK14Eym8Z38buSfCC", "5f6e0146e64d69ae9aa929604467d965f63f3f473a0f8fa6fb5804878f72b4da"})
	list = append(list, &Account{"TASdSCbgDeDREmF5F6bR4kUVidVG7PFKT6", "0eef3783214f486d824c0bc7938860c82e6f1be5bf0a90f60d67b0ba7a4c379e"})
	list = append(list, &Account{"TBtdY3F6573ttBBA1Gh8EoKCyteaUj3LKW", "9c0e43137ec44ca3a6995fc571fc04c62e0245fc3d3eb95775ae8ba69bab866d"})

	toAddress := "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	amount := int64(30000 * 1e6)
	for index, item := range list {
		if err := TransferTrx(item.Owner, item.Key, toAddress, amount); err != nil {
			log.Errorf(err, "%v, TransferTrx error, owner:%v", index, item.Owner)
		} else {
			log.Infof("%v, TransferTrx success, owner:%v", index, item.Owner)
		}
	}

}

type Account struct {
	Owner string
	Key   string
}
