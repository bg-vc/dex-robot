package service

import "github.com/vincecfl/go-common/log"

func TransferTrxHandle() {
	list := make([]*Account, 0)

	// 0
	list = append(list, &Account{"TPMXwJuAt6Nbi5M3njhhUsEHw5QUCx5VV2", "1b93c7501b89d1e239fcfac50270036863271b31736c7960984a6be2c02c902a"})
	list = append(list, &Account{"TLcLcDTDdsKc1ZRDFgt6JPKYyquX9v25V4", "4fcf0b8df38091c9908ee52f75edaa9bf16d55c9ade992dd043a8fb684efa663"})

	// 1
	list = append(list, &Account{"TKeHVce6sxDdG7SbNfN81v1PD9nbTGeSXH", "27c8c9144a576820e35d2ecd3b277a9b264430d90903301376c3aa914db4ca14"})
	list = append(list, &Account{"TPfaaP5WYt5L8TLtkVYzQXVvLeVPHqR8cX", "e873daf5b5549493bb38299e404e6d3442d814e5aad6d4d27d043190e02ca814"})
	list = append(list, &Account{"TTuwCHR9ofuV6k2wheBLf4wDEaMcyM5Dk7", "785473388246e048645488f52eb5fcaa2b80f0a23e3f33920b04e23cb98351d5"})

	// 2
	list = append(list, &Account{"TBFpBBrAgKvSm1cRK9EfwPXCupSbqH3yXB", "A95D7AC43F0360FCDE2F77380D492683A38E474B87E71088503DFE1E6F4473FB"})
	list = append(list, &Account{"TXULv92s3eMMpeyFoUveVsAwqbKqMrzW1x", "BA04FB6150F4E2A3FE6FE383D1A782E6540B4AA9C89ABB9E68DE536B7E68CE21"})
	list = append(list, &Account{"TDkyCoEo2XG3HFa21b4YGC4szqBaHoDHjj", "2294BA50616CDF6065E055BE81F06C12651BD53F71E57D7164F11E05AE1EDAAB"})

	// 3
	list = append(list, &Account{"TTCT6rgLXGnoJwBuMrVqXjRyyAFgdtgHro", "A6662FE759D9FF76143D5635D9D607D148A3604A5C64C4E34454D6BE9215E3C5"})
	list = append(list, &Account{"TJmmJEfCGvZMYe9FSZ2niVFX1omx43HnzK", "A45876643F9C7F573F29E7FA34294637A1DCA210B4565D69219A9EBAF11C5549"})
	list = append(list, &Account{"TUkWgBw178XuimVCkSPqp2Mjyrfpp9ApKt", "5B6AAB5711FC4C27F8C970708CE38A823D76EF9996E8BBDE9547B17354871849"})

	// 4
	list = append(list, &Account{"TPwqzXcUgQaFTFMU6AefcbKedB2u1D1qiV", "427CAFECEB52553FBF3B31119633E26408AE884323C3EDB13A71FFC82E4346B0"})
	list = append(list, &Account{"TSxtvS7DTmZGvm3nnxmu6DysCZ8a1JtG7E", "53B3AB7ADAB16110A37450F6F19DEA893C5A2D31DDED098488E95CF0ACB27EEE"})
	list = append(list, &Account{"TXGSMzBT9RuW12t3rdP7tc615UBmGRWBXS", "3BF6E7522C01515D688FC117BBD20371A6C4D2A8A4651D70D69D6505C2D48F94"})

	// 5
	list = append(list, &Account{"TDo2Sg79RBuri8j839M5a3N3M8VkVqPT64", "5CCF93D74CE566B3857267C4A92CBDE46A5C68E84A0BC468DC345825345C85A2"})
	list = append(list, &Account{"TRdTi9rTKAdaWDKZRU9aeBV9ii91gQ3KKs", "DC66A127B9150CE1821253F0DB8B3F2E0901913BFF255B643C5AE079D1DA5B03"})
	list = append(list, &Account{"TPvB6nRgcvTxvMxs3yxi7UQQjh2ex4UoYD", "508CC008EC2CB5B98C742FB30326D41AB2CA8A4A3EDCC4CE7EBEB3457E3529CA"})

	toAddress := "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	amount := int64(30000 * 1e6)
	for _, item := range list {
		if err := TransferTrx(item.Owner, item.Key, toAddress, amount); err != nil {
			log.Errorf(err, "TransferTrx error, owner:%v", item.Owner)
		} else {
			log.Infof("TransferTrx success, owner:%v", item.Owner)
		}
	}

}

type Account struct {
	Owner string
	Key   string
}
