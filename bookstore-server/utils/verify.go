package utils

// 验证规则
var (
	BooksCategoryVerify           = Rules{"CategoryRank": {NotEmpty()}, "CategoryName": {NotEmpty()}}
	AdminUserRegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	MallUserRegisterVerify        = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	AdminUserChangePasswordVerify = Rules{"Password": {NotEmpty()}}
	BooksAddParamVerify           = Rules{"BooksName": {Le("128")}, "BooksIntro": {Le("200")}, "BooksCategoryId": {Ge("1")}, "BooksCoverImg": {NotEmpty()}, "OriginalPrice": {Ge("0"), Le("1000000")},
		"sellingPrice": {Ge("1"), Le("1000000")}, "stockNum": {Ge("1"), Le("100000")}, "Tag": {Le("16")}, "booksDetailContent": {NotEmpty()}}
	CarouselAddParamVerify       = Rules{"CarouselUrl": {NotEmpty()}, "RedirectUrl": {NotEmpty()}, "CarouselRank": {NotEmpty(), Ge("0"), Le("200")}}
	IndexConfigAddParamVerify    = Rules{"ConfigName": {NotEmpty()}, "ConfigType": {Ge("1"), Le("5")}, "BooksId": {NotEmpty()}, "ConfigRank": {Ge("1"), Le("1000000")}}
	IndexConfigUpdateParamVerify = Rules{"ConfigId": {NotEmpty()}, "ConfigName": {NotEmpty()}, "ConfigType": {Ge("1"), Le("5")}, "BooksId": {NotEmpty()}, "ConfigRank": {Ge("1"), Le("1000000")}}
	SaveOrderParamVerify         = Rules{"CartItemIds": {NotEmpty()}, "AddressId": {NotEmpty()}}
)
