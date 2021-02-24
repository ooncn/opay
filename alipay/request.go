package alipay

import (
	"errors"
	"github.com/ooncn/common/util"
	"time"
)

type AftAifinFireeyeOcrImageQuery struct {
}
type AftFinsecureRiskplusSecurityPolicyQuery struct {
}
type AliosOpenAutoInfoQuery struct {
}
type AlipayAccountExrateAdviceAccept struct {
}
type AlipayAccountExrateAllclientrateQuery struct {
}
type AlipayAccountExrateRatequery struct {
}
type AlipayAccountExrateTraderequestCreate struct {
}
type AlipayAcquireCancel struct {
}
type AlipayAcquireClose struct {
}
type AlipayAcquireCreateandpay struct {
}
type AlipayAcquirePrecreate struct {
}
type AlipayAcquireQuery struct {
}
type AlipayAcquireRefund struct {
}
type AlipayAppTokenGet struct {
}
type AlipayAssetPointBalanceQuery struct {
}
type AlipayAssetPointBudgetQuery struct {
}
type AlipayAssetPointOrderCreate struct {
}
type AlipayAssetPointOrderQuery struct {
}
type AlipayBossCsChannelQuery struct {
}
type AlipayBossFncXwbtestRetModify struct {
}
type AlipayBossProdArrangementOfflineQuery struct {
}
type AlipayBossProdMyTestQuery struct {
}
type AlipayCommerceAdContractSign struct {
}
type AlipayCommerceAirCallcenterTradeApply struct {
}
type AlipayCommerceAirXfgDsgModify struct {
}
type AlipayCommerceCityfacilitatorCityQuery struct {
}
type AlipayCommerceCityfacilitatorDepositCancel struct {
}
type AlipayCommerceCityfacilitatorDepositConfirm struct {
}
type AlipayCommerceCityfacilitatorDepositQuery struct {
}
type AlipayCommerceCityfacilitatorFunctionQuery struct {
}
type AlipayCommerceCityfacilitatorScriptQuery struct {
}
type AlipayCommerceCityfacilitatorStationQuery struct {
}
type AlipayCommerceCityfacilitatorVoucherBatchquery struct {
}
type AlipayCommerceCityfacilitatorVoucherCancel struct {
}
type AlipayCommerceCityfacilitatorVoucherConfirm struct {
}
type AlipayCommerceCityfacilitatorVoucherGenerate struct {
}
type AlipayCommerceCityfacilitatorVoucherQuery struct {
}
type AlipayCommerceCityfacilitatorVoucherRefund struct {
}
type AlipayCommerceCityfacilitatorVoucherUpload struct {
}
type AlipayCommerceDataCampaignCreate struct {
}
type AlipayCommerceDataCampaignSend struct {
}
type AlipayCommerceDataCustommetricSync struct {
}
type AlipayCommerceDataLogdataSync struct {
}
type AlipayCommerceDataMonitordataSync struct {
}
type AlipayCommerceEducateAuthenticateCampuscardCreate struct {
}
type AlipayCommerceEducateParttimejobInfoCreate struct {
}
type AlipayCommerceEducateSchoolcodeTokenCreate struct {
}
type AlipayCommerceEducateStudentinfoShare struct {
}
type AlipayCommerceEducateUserClickCreate struct {
}
type AlipayCommerceIotDeviceserviceCancel struct {
}
type AlipayCommerceLogisticsWaybillMinimctSync struct {
}
type AlipayCommerceLotteryPresentlistQuery struct {
}
type AlipayCommerceLotteryPresentSend struct {
}
type AlipayCommerceLotteryTypelistQuery struct {
}
type AlipayCommerceTransportIntelligentizeDataSync struct {
}
type AlipayCommerceTransportNfccardSend struct {
}
type AlipayCommerceTransportOfflinepayKeyQuery struct {
}
type AlipayCommerceTransportOfflinepayRecordVerify struct {
}
type AlipayCommerceTransportOfflinepayUserblacklistQuery struct {
}
type AlipayCommerceTransportParkingReserveConfirm struct {
}
type AlipayCommerceTransportVehicleownerMessageSend struct {
}
type AlipayDaoweiOrderCancel struct {
}
type AlipayDaoweiOrderConfirm struct {
}
type AlipayDaoweiOrderModify struct {
}
type AlipayDaoweiOrderQuery struct {
}
type AlipayDaoweiOrderRefund struct {
}
type AlipayDaoweiOrderRefuse struct {
}
type AlipayDaoweiOrderSpModify struct {
}
type AlipayDaoweiOrderTransfer struct {
}
type AlipayDaoweiServiceModify struct {
}
type AlipayDaoweiSpModify struct {
}
type AlipayDaoweiSpScheduleModify struct {
}
type AlipayDataAiserviceSgxGatewayQuery struct {
}
type AlipayDataAiserviceSmartpriceMerchanteffectQuery struct {
}
type AlipayDataBillAccountlogQuery struct {
}
type AlipayDataBillBailQuery struct {
}
type AlipayDataBillBalancehisQuery struct {
}
type AlipayDataBillBalanceQuery struct {
}
type AlipayDataBillBuyQuery struct {
}
type AlipayDataBillDownloadurlGet struct {
}
type AlipayDataBillEreceiptApply struct {
}
type AlipayDataBillEreceiptQuery struct {
}
type AlipayDataBillSellQuery struct {
}
type AlipayDataBillTransferQuery struct {
}
type AlipayDataDataexchangeSfasdf struct {
}
type AlipayDataDataserviceAdCreativeBatchquery struct {
}
type AlipayDataDataserviceAdCreativeCreateormodify struct {
}
type AlipayDataDataserviceAdCreativeQuery struct {
}
type AlipayDataDataserviceAdDataQuery struct {
}
type AlipayDataDataserviceAdGroupCreateormodify struct {
}
type AlipayDataDataserviceAdGroupQuery struct {
}
type AlipayDataDataserviceAdOffline struct {
}
type AlipayDataDataserviceAdOnline struct {
}
type AlipayDataDataserviceAdPlanCreateormodify struct {
}
type AlipayDataDataserviceAdPlanQuery struct {
}
type AlipayDataDataserviceAdPrincipalConsult struct {
}
type AlipayDataDataserviceAdPrincipalCreateormodify struct {
}
type AlipayDataDataserviceAdPrincipalQuery struct {
}
type AlipayDataDataserviceAdPromotepageQuery struct {
}
type AlipayDataDataserviceAdUserbalanceOffline struct {
}
type AlipayDataDataserviceAdUserbalanceOnline struct {
}
type AlipayDataDataserviceAdUserCreate struct {
}
type AlipayDataDataserviceBillDownloadurlQuery struct {
}
type AlipayDataDataserviceChinaremodelQuery struct {
}
type AlipayDataDataserviceCodeReco struct {
}
type AlipayDataDataserviceSdfsdf struct {
}
type AlipayDataDataserviceShoppingmallrecShopQuery struct {
}
type AlipayDataDataserviceShoppingmallrecVoucherQuery struct {
}
type AlipayDataDataserviceUserlevelZrankGet struct {
}
type AlipayEbppBillAdd struct {
}
type AlipayEbppBillGet struct {
}
type AlipayEbppBillSearch struct {
}
type AlipayEbppInvoiceApplyResultSync struct {
}
type AlipayEbppInvoiceAuthSign struct {
}
type AlipayEbppInvoiceAuthUnsign struct {
}
type AlipayEbppInvoiceInfoSend struct {
}
type AlipayEbppInvoiceMerchantEnterstatusQuery struct {
}
type AlipayEbppInvoiceMerchantlistEnterApply struct {
}
type AlipayEbppInvoiceSyncSimpleSend struct {
}
type AlipayEbppInvoiceTitleDynamicGet struct {
}
type AlipayEbppInvoiceTitleListGet struct {
}
type AlipayEbppMerchantConfigGet struct {
}
type AlipayEbppPdeductBillPayStatus struct {
}
type AlipayEbppPdeductPay struct {
}
type AlipayEbppPdeductSignAdd struct {
}
type AlipayEbppPdeductSignCancel struct {
}
type AlipayEbppPdeductSignQuery struct {
}
type AlipayEbppPdeductSignValidate struct {
}
type AlipayEcapiprodCreditGet struct {
}
type AlipayEcapiprodDataPut struct {
}
type AlipayEcapiprodDrawndnContractGet struct {
}
type AlipayEcapiprodDrawndnDrawndnlistQuery struct {
}
type AlipayEcapiprodDrawndnFeerecordQuery struct {
}
type AlipayEcapiprodDrawndnLendingrecordQuery struct {
}
type AlipayEcapiprodDrawndnPaymentscheduleGet struct {
}
type AlipayEcapiprodDrawndnRepaymentrecordQuery struct {
}
type AlipayEcardEduPublicBind struct {
}
type AlipayEcoCityserviceMessageSend struct {
}
type AlipayEcoContractSignflowsCreate struct {
}
type AlipayEcoCplifeBasicserviceInitialize struct {
}
type AlipayEcoCplifeBasicserviceModify struct {
}
type AlipayEcoCplifeBillBatchquery struct {
}
type AlipayEcoCplifeBillBatchUpload struct {
}
type AlipayEcoCplifeBillDelete struct {
}
type AlipayEcoCplifeBillModify struct {
}
type AlipayEcoCplifeBillSync struct {
}
type AlipayEcoCplifeCommunityBatchquery struct {
}
type AlipayEcoCplifeCommunityCreate struct {
}
type AlipayEcoCplifeCommunityDetailsQuery struct {
}
type AlipayEcoCplifeCommunityModify struct {
}
type AlipayEcoCplifeNoticeDelete struct {
}
type AlipayEcoCplifeNoticePublish struct {
}
type AlipayEcoCplifePayResultQuery struct {
}
type AlipayEcoCplifeRepairStatusUpdate struct {
}
type AlipayEcoCplifeResidentinfoDelete struct {
}
type AlipayEcoCplifeResidentinfoUpload struct {
}
type AlipayEcoCplifeRoominfoDelete struct {
}
type AlipayEcoCplifeRoominfoQuery struct {
}
type AlipayEcoCplifeRoominfoUpload struct {
}
type AlipayEcoCplifeRooominfoQuery struct {
}
type AlipayEcoCplifeUseridentityStatusUpdate struct {
}
type AlipayEcoDocTemplateCreate struct {
}
type AlipayEcoDoctemplateSettingurlQuery struct {
}
type AlipayEcoEduKtBillingModify struct {
}
type AlipayEcoEduKtBillingQuery struct {
}
type AlipayEcoEduKtBillingSend struct {
}
type AlipayEcoEduKtParentQuery struct {
}
type AlipayEcoEduKtSchoolinfoModify struct {
}
type AlipayEcoEduKtStudentModify struct {
}
type AlipayEcoEduKtStudentQuery struct {
}
type AlipayEcoFilePathQuery struct {
}
type AlipayEcoMycarCarlibInfoPush struct {
}
type AlipayEcoMycarCarmodelModify struct {
}
type AlipayEcoMycarDataExternalQuery struct {
}
type AlipayEcoMycarDataExternalSend struct {
}
type AlipayEcoMycarDataserviceViolationinfoShare struct {
}
type AlipayEcoMycarMaintainDataUpdate struct {
}
type AlipayEcoMycarMaintainOrderCreate struct {
}
type AlipayEcoMycarMaintainOrderstatusUpdate struct {
}
type AlipayEcoMycarOrderStatusQuery struct {
}
type AlipayEcoMycarParkingCardbarcodeCreate struct {
}
type AlipayEcoMycarParkingChargeinfoSync struct {
}
type AlipayEcoMycarParkingConfigQuery struct {
}
type AlipayEcoMycarParkingConfigSet struct {
}
type AlipayEcoMycarParkingEnterinfoSync struct {
}
type AlipayEcoMycarParkingExitinfoSync struct {
}
type AlipayEcoMycarParkingLotbarcodeCreate struct {
}
type AlipayEcoMycarParkingOrderstatusQuery struct {
}
type AlipayEcoMycarParkingOrderSync struct {
}
type AlipayEcoMycarParkingOrderUpdate struct {
}
type AlipayEcoMycarParkingParkinglotinfoCreate struct {
}
type AlipayEcoMycarParkingParkinglotinfoQuery struct {
}
type AlipayEcoMycarParkingParkinglotinfoUpdate struct {
}
type AlipayEcoMycarParkingSpaceinfoSync struct {
}
type AlipayEcoMycarParkingVehicleQuery struct {
}
type AlipayEcoMycarPromoTicketPush struct {
}
type AlipayEcoMycarPromoTicketSync struct {
}
type AlipayEcoMycarPromoVoucherVerify struct {
}
type AlipayEcoMycarTradeRefund struct {
}
type AlipayEcoMycarViolationCityPush struct {
}
type AlipayEcoMycarViolationInfoPush struct {
}
type AlipayEcoSignFlowCancel struct {
}
type AlipayEcoSignFlowCreate struct {
}
type AlipayEcoSignFlowQuery struct {
}
type AlipayEcoSignflowsDetailQuery struct {
}
type AlipayEcoSignflowsUrlQuery struct {
}
type AlipayEcoWelfareCodeSync struct {
}
type AlipayExscUserCurrentsignGet struct {
}
type AlipayExscUserFirstfundinpourGet struct {
}
type AlipayExscUserFirstsignGet struct {
}
type AlipayFlashsalesStockSyncUpdate struct {
}
type AlipayFundAccountbookCreate struct {
}
type AlipayFundAccountbookQuery struct {
}
type AlipayFundAccountQuery struct {
}
type AlipayFundAuthOperationCancel struct {
}
type AlipayFundAuthOperationDetailQuery struct {
}
type AlipayFundAuthOrderAppFreeze struct {
}
type AlipayFundAuthOrderFreeze struct {
}
type AlipayFundAuthOrderUnfreeze struct {
}
type AlipayFundAuthOrderVoucherCreate struct {
}
type AlipayFundBatchClose struct {
}
type AlipayFundBatchCreate struct {
}
type AlipayFundBatchDetailQuery struct {
}
type AlipayFundBatchTransfer struct {
}
type AlipayFundBatchUniTransfer struct {
}
type AlipayFundCouponOperationQuery struct {
}
type AlipayFundCouponOrderAgreementPay struct {
}
type AlipayFundCouponOrderAppPay struct {
}
type AlipayFundCouponOrderDisburse struct {
}
type AlipayFundCouponOrderPagePay struct {
}
type AlipayFundCouponOrderRefund struct {
}
type AlipayFundTransAppPay struct {
}
type AlipayFundTransCommonQuery struct {
}
type AlipayFundTransOrderQuery struct {
}
type AlipayFundTransRefund struct {
}
type AlipayFundTransToaccountTransfer struct {
}
type AlipayFundTransUniTransfer struct {
}
type AlipayInsAutoAutoaftermarketAttachmentUpload struct {
}
type AlipayInsAutoAutoaftermarketDepotCreateormodify struct {
}
type AlipayInsAutoAutoaftermarketInserviceorderNotify struct {
}
type AlipayInsAutoAutoaftermarketOutorderSync struct {
}
type AlipayInsAutoAutoinsprodCommonConsult struct {
}
type AlipayInsAutoAutoinsprodEnquriyApply struct {
}
type AlipayInsAutoAutoinsprodPolicyApply struct {
}
type AlipayInsAutoAutoinsprodPolicyCancel struct {
}
type AlipayInsAutoAutoinsprodQuoteApply struct {
}
type AlipayInsAutoAutoinsprodQuoteQuery struct {
}
type AlipayInsAutoAutoinsprodUserCertify struct {
}
type AlipayInsAutoCarSave struct {
}
type AlipayInsCooperationProductOfflineBatchquery struct {
}
type AlipayInsCooperationProductQrcodeApply struct {
}
type AlipayInsCooperationRegionQrcodeApply struct {
}
type AlipayInsSceneApplicationIssueConfirm struct {
}
type AlipayInsSceneCouponReceive struct {
}
type AlipayInsSceneCouponSend struct {
}
type AlipayMarketingCampaignActivityOfflineCreate struct {
}
type AlipayMarketingCampaignActivityOfflineTrigger struct {
}
type AlipayMarketingCampaignCashCreate struct {
}
type AlipayMarketingCampaignCashDetailQuery struct {
}
type AlipayMarketingCampaignCashListQuery struct {
}
type AlipayMarketingCampaignCashStatusModify struct {
}
type AlipayMarketingCampaignCashTrigger struct {
}
type AlipayMarketingCampaignCertCreate struct {
}
type AlipayMarketingCampaignDiscountBudgetAppend struct {
}
type AlipayMarketingCampaignDiscountBudgetCreate struct {
}
type AlipayMarketingCampaignDiscountBudgetQuery struct {
}
type AlipayMarketingCampaignDiscountQuery struct {
}
type AlipayMarketingCampaignDiscountStatusUpdate struct {
}
type AlipayMarketingCampaignDiscountWhitelistQuery struct {
}
type AlipayMarketingCampaignDiscountWhitelistUpdate struct {
}
type AlipayMarketingCampaignDrawcampCreate struct {
}
type AlipayMarketingCampaignDrawcampQuery struct {
}
type AlipayMarketingCampaignDrawcampStatusUpdate struct {
}
type AlipayMarketingCampaignDrawcampUpdate struct {
}
type AlipayMarketingCampaignDrawcampWhitelistCreate struct {
}
type AlipayMarketingCampaignPrizeAmountQuery struct {
}
type AlipayMarketingCardActivateformQuery struct {
}
type AlipayMarketingCardActivateurlApply struct {
}
type AlipayMarketingCardBenefitCreate struct {
}
type AlipayMarketingCardBenefitDelete struct {
}
type AlipayMarketingCardBenefitModify struct {
}
type AlipayMarketingCardBenefitQuery struct {
}
type AlipayMarketingCardConsumeSync struct {
}
type AlipayMarketingCardDelete struct {
}
type AlipayMarketingCardFormtemplateSet struct {
}
type AlipayMarketingCardOpen struct {
}
type AlipayMarketingCardQuery struct {
}
type AlipayMarketingCardTemplateBatchquery struct {
}
type AlipayMarketingCardTemplateCreate struct {
}
type AlipayMarketingCardTemplateModify struct {
}
type AlipayMarketingCardTemplateQuery struct {
}
type AlipayMarketingCardUpdate struct {
}
type AlipayMarketingCashitemvoucherTemplateCreate struct {
}
type AlipayMarketingCashlessitemvoucherTemplateCreate struct {
}
type AlipayMarketingCashlessvoucherTemplateCreate struct {
}
type AlipayMarketingCashlessvoucherTemplateModify struct {
}
type AlipayMarketingCashvoucherTemplateCreate struct {
}
type AlipayMarketingCashvoucherTemplateModify struct {
}
type AlipayMarketingCdpAdvertiseCreate struct {
}
type AlipayMarketingCdpAdvertiseModify struct {
}
type AlipayMarketingCdpAdvertiseOperate struct {
}
type AlipayMarketingCdpAdvertiseQuery struct {
}
type AlipayMarketingCdpAdvertiseReportQuery struct {
}
type AlipayMarketingCdpRecommendQuery struct {
}
type AlipayMarketingCouponTemplateCreate struct {
}
type AlipayMarketingExchangevoucherTemplateCreate struct {
}
type AlipayMarketingExchangevoucherUse struct {
}
type AlipayMarketingFacetofaceDecodeUse struct {
}
type AlipayMarketingMaterialImageUpload struct {
}
type AlipayMarketingToolFengdieActivityCreate struct {
}
type AlipayMarketingToolFengdieActivityQuery struct {
}
type AlipayMarketingToolFengdieEditorQuery struct {
}
type AlipayMarketingToolFengdieMemberCreate struct {
}
type AlipayMarketingToolFengdieSitesBatchquery struct {
}
type AlipayMarketingToolFengdieSitesConfirm struct {
}
type AlipayMarketingToolFengdieSitesCreate struct {
}
type AlipayMarketingToolFengdieSitesDelete struct {
}
type AlipayMarketingToolFengdieSitesQuery struct {
}
type AlipayMarketingToolFengdieSitesSync struct {
}
type AlipayMarketingToolFengdieSpaceBatchquery struct {
}
type AlipayMarketingToolFengdieSpaceCreate struct {
}
type AlipayMarketingToolFengdieSpaceQuery struct {
}
type AlipayMarketingToolFengdieTemplateBatchquery struct {
}
type AlipayMarketingToolFengdieTemplateQuery struct {
}
type AlipayMarketingToolFengdieTemplateSend struct {
}
type AlipayMarketingUserulePidQuery struct {
}
type AlipayMarketingVoucherAuthSend struct {
}
type AlipayMarketingVoucherConfirm struct {
}
type AlipayMarketingVoucherListQuery struct {
}
type AlipayMarketingVoucherQuery struct {
}
type AlipayMarketingVoucherSend struct {
}
type AlipayMarketingVoucherStockCreate struct {
}
type AlipayMarketingVoucherStockMatch struct {
}
type AlipayMarketingVoucherStockQuery struct {
}
type AlipayMarketingVoucherStockUse struct {
}
type AlipayMarketingVoucherTemplateDelete struct {
}
type AlipayMarketingVoucherTemplatedetailQuery struct {
}
type AlipayMarketingVoucherTemplatelistQuery struct {
}
type AlipayMdataTagGet struct {
}
type AlipayMemberCouponQuerylist struct {
}
type AlipayMerchantItemFileUpload struct {
}
type AlipayMerchantOrderConsumerQuery struct {
}
type AlipayMerchantOrderDigestConsumerBatchquery struct {
}
type AlipayMerchantOrderSecuritydetailConsumerQuery struct {
}
type AlipayMerchantOrderSecuritydigestConsumerBatchquery struct {
}
type AlipayMerchantOrderSync struct {
}
type AlipayMicropayOrderConfirmpayurlGet struct {
}
type AlipayMicropayOrderDirectPay struct {
}
type AlipayMicropayOrderFreezepayurlGet struct {
}
type AlipayMicropayOrderFreeze struct {
}
type AlipayMicropayOrderGet struct {
}
type AlipayMicropayOrderUnfreeze struct {
}
type AlipayMobileBeaconDeviceAdd struct {
}
type AlipayMobileBeaconDeviceDelete struct {
}
type AlipayMobileBeaconDeviceModify struct {
}
type AlipayMobileBeaconDeviceQuery struct {
}
type AlipayMobileBeaconMessageSend struct {
}
type AlipayMobileBksigntokenVerify struct {
}
type AlipayMobileCodeCreate struct {
}
type AlipayMobileCodeQuery struct {
}
type AlipayMobilePublicAccountAdd struct {
}
type AlipayMobilePublicAccountDelete struct {
}
type AlipayMobilePublicAccountQuery struct {
}
type AlipayMobilePublicAccountReset struct {
}
type AlipayMobilePublicAppinfoUpdate struct {
}
type AlipayMobilePublicContactFollowList struct {
}
type AlipayMobilePublicFollowList struct {
}
type AlipayMobilePublicGisGet struct {
}
type AlipayMobilePublicInfoModify struct {
}
type AlipayMobilePublicInfoQuery struct {
}
type AlipayMobilePublicLabelAdd struct {
}
type AlipayMobilePublicLabelDelete struct {
}
type AlipayMobilePublicLabelQuery struct {
}
type AlipayMobilePublicLabelUpdate struct {
}
type AlipayMobilePublicLabelUserAdd struct {
}
type AlipayMobilePublicLabelUserDelete struct {
}
type AlipayMobilePublicLabelUserQuery struct {
}
type AlipayMobilePublicMenuAdd struct {
}
type AlipayMobilePublicMenuDelete struct {
}
type AlipayMobilePublicMenuGet struct {
}
type AlipayMobilePublicMenuQuery struct {
}
type AlipayMobilePublicMenuUpdate struct {
}
type AlipayMobilePublicMenuUserQuery struct {
}
type AlipayMobilePublicMenuUserUpdate struct {
}
type AlipayMobilePublicMessagebatchPush struct {
}
type AlipayMobilePublicMessageCustomSend struct {
}
type AlipayMobilePublicMessageLabelSend struct {
}
type AlipayMobilePublicMessagePush struct {
}
type AlipayMobilePublicMessageSingleSend struct {
}
type AlipayMobilePublicMessagespecifyPush struct {
}
type AlipayMobilePublicMessageTotalSend struct {
}

//为了满足用户渠道推广分析的需要，公众平台提供了生成带参数二维码的接口。使用该接口可以获得多个带不同场景值的二维码，用户扫描后，公众号可以接收到事件推送。目前有2种类型的二维码，分别是临时二维码、和永久二维码，前者有过期时间，最大为1800秒。
//每次创建二维码ticket需要提供一个开发者自行设定的参数（scene_id）。
type AlipayMobilePublicQrcodeCreate struct {
	AlipayMobilePublicQrcodeCreateResponse *AlipayMobilePublicQrcodeCreateResponse `json:"alipay_mobile_public_qrcode_create_response"`
	ErrorResponse                          *ErrorResponse                          `json:"error_response"`
	Sign                                   string                                  `json:"sign"`
	/*
		{
		    "codeInfo": { //开发者自定义信息
		        "scene": {
		            "sceneId": "1234"//场景Id，最长32位，英文字母、数字以及下划线，开发者自定义
		        },
		        "gotoUrl": "http://www.******.com"//跳转URL，扫码关注服务窗后会直接跳转到此URL
		    },
		    "codeType": "TEMP",//二维码类型，目前只支持两种类型：TEMP：临时的（默认）、PERM：永久的
		    "expireSecond": "1800",//临时码过期时间，以秒为单位，最大不超过1800秒；永久码置空
		    "showLogo": "Y"//二维码中间是否显示服务窗logo，Y：显示；N：不显示（默认）
		}
	*/
} //alipay.mobile.public.qrcode.create 生成带参数的二维码
type AlipayMobilePublicQrcodeCreateResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
} //alipay.mobile.public.qrcode.create 生成带参数的二维码
type AlipayMobilePublicShortlinkCreate struct {
}
type AlipayMobilePublicTemplateMessageDelete struct {
}
type AlipayMobilePublicTemplateMessageGet struct {
}
type AlipayMobilePublicTemplateMessageModify struct {
}
type AlipayMobilePublicTemplateMessageQuery struct {
}
type AlipayMobileRecommendGet struct {
}
type AlipayMobileShakeUserQuery struct {
}
type AlipayMobileStdPublicAccountQuery struct {
}
type AlipayMobileStdPublicExpressUserQuery struct {
}
type AlipayMobileStdPublicFollowList struct {
}
type AlipayMobileStdPublicMenuQuery struct {
}
type AlipayMobileStdPublicMessageCustomSend struct {
}
type AlipayMpointprodBenefitDetailGet struct {
}
type AlipayMsaasMediarecogVoiceMediaaudioUpload struct {
}
type AlipayMsaasPromotionCpainfoCreate struct {
}
type AlipayOfflineMarketApplyorderBatchquery struct {
}
type AlipayOfflineMarketingVoucherCodeUpload struct {
}
type AlipayOfflineMarketingVoucherCreate struct {
}
type AlipayOfflineMarketingVoucherModify struct {
}
type AlipayOfflineMarketingVoucherOffline struct {
}
type AlipayOfflineMarketingVoucherStatusQuery struct {
}
type AlipayOfflineMarketingVoucherUse struct {
}
type AlipayOfflineMarketItemCreate struct {
}
type AlipayOfflineMarketItemModify struct {
}
type AlipayOfflineMarketItemState struct {
}
type AlipayOfflineMarketMcommentQuery struct {
}
type AlipayOfflineMarketProductBatchquery struct {
}
type AlipayOfflineMarketProductQuerydetail struct {
}
type AlipayOfflineMarketReporterrorCreate struct {
}
type AlipayOfflineMarketShopApplyorderCancel struct {
}
type AlipayOfflineMarketShopBatchquery struct {
}
type AlipayOfflineMarketShopCategoryQuery struct {
}
type AlipayOfflineMarketShopCreate struct {
}
type AlipayOfflineMarketShopDiscountQuery struct {
}
type AlipayOfflineMarketShopModify struct {
}
type AlipayOfflineMarketShopPublicBind struct {
}
type AlipayOfflineMarketShopPublicUnbind struct {
}
type AlipayOfflineMarketShopQuerydetail struct {
}
type AlipayOfflineMarketShopSummaryBatchquery struct {
}
type AlipayOfflineMaterialImageDownload struct {
}
type AlipayOfflineMaterialImageUpload struct {
}
type AlipayOfflineProviderDishQuery struct {
}
type AlipayOfflineProviderEquipmentAuthQuerybypage struct {
}
type AlipayOfflineProviderEquipmentAuthRemove struct {
}
type AlipayOfflineProviderMonitorLogSync struct {
}
type AlipayOfflineProviderShopactionRecord struct {
}
type AlipayOfflineProviderUseractionRecord struct {
}
type AlipayOpenAgentCancel struct {
}
type AlipayOpenAgentConfirm struct {
}
type AlipayOpenAgentCreate struct {
}
type AlipayOpenAgentFacetofaceSign struct {
}
type AlipayOpenAgentMiniCreate struct {
}
type AlipayOpenAgentMobilepaySign struct {
}
type AlipayOpenAgentOfflinepaymentSign struct {
}
type AlipayOpenAgentOrderQuery struct {
}
type AlipayOpenAgentSignstatusQuery struct {
}
type AlipayOpenAgentZhimabriefSign struct {
}
type AlipayOpenAppAlipaycertDownload struct {
}
type AlipayOpenAppCallQuery struct {
}
type AlipayOpenAppCodetesttest struct {
}
type AlipayOpenAppDedfDdQuery struct {
}
type AlipayOpenAppDfsfasDeQuery struct {
}
type AlipayOpenAppLingbalingliuQuery struct {
}
type AlipayOpenAppLingjiuyisiCreate struct {
}
type AlipayOpenAppLingjiuyiwuBatchquery struct {
}
type AlipayOpenAppMembersCreate struct {
}
type AlipayOpenAppMembersDelete struct {
}
type AlipayOpenAppMembersQuery struct {
}
type AlipayOpenAppMessageSubscriptionModify struct {
}
type AlipayOpenAppMessageSubscriptionQuery struct {
}
type AlipayOpenAppMessageTopicSubscribe struct {
}
type AlipayOpenAppMessageTopicUnsubscribe struct {
}
type AlipayOpenAppMiniTemplatemessageSend struct {
	ToUserId       string `json:"to_user_id"`       //触达消息的支付宝user_id
	FormId         string `json:"form_id"`          //用户发生的交易行为的交易号
	UserTemplateId string `json:"user_template_id"` //用户申请的模板 ID号
	Page           string `json:"page"`             //小程序的跳转页面
	Data           string `json:"data"`             //开发者需要发送模板消息中的自定义部分来替换模板的占位符{“keyword1”: {“value” : “12:00”}, “keyword2”: {“value” : “20180808”}, “keyword3”: {“value” : “支付宝”}}
} // 模板消息
type AlipayOpenAppNotifyModify struct {
}
type AlipayOpenAppNotifyVerify struct {
}
type AlipayOpenAppOpenbizmockApisdkgrayQuery struct {
}
type AlipayOpenAppOpenbizmockMessageSend struct {
}
type AlipayOpenAppQrcodeCreate struct {
}
type AlipayOpenAppSmgMsgSend struct {
}
type AlipayOpenAppSmsgDataSync struct {
}
type AlipayOpenAppUpdattestBatchquery struct {
}
type AlipayOpenAppXwbtestabcQuery struct {
}
type AlipayOpenAppYiyiyiwuQuery struct {
}
type AlipayOpenAppYufanlingsanyaowuYufalingsanyaowuQuery struct {
}
type AlipayOpenAuthAppAesGet struct {
}
type AlipayOpenAuthAppAesSet struct {
}
type AlipayOpenAuthIndustryPlatformCreateToken struct {
}
type AlipayOpenAuthTokenAppQuery struct {
}
type AlipayOpenAuthTokenApp struct {
}
type AlipayOpenBizCreate struct {
}
type AlipayOpenDafBatchquery struct {
}
type AlipayOpenEchoOfflineSend struct {
}
type AlipayOpenInviteOrderCreate struct {
}
type AlipayOpenInviteOrderQuery struct {
}
type AlipayOpenMessagetestCesSend struct {
}
type AlipayOpenMiniBaseinfoModify struct {
}
type AlipayOpenMiniBaseinfoQuery struct {
}
type AlipayOpenMiniBizdataTemplatemessageDelete struct {
}
type AlipayOpenMiniBizdataTemplatemessageUpload struct {
}
type AlipayOpenMiniCategoryQuery struct {
}
type AlipayOpenMiniContentSync struct {
}
type AlipayOpenMiniDataVisitQuery struct {
}
type AlipayOpenMiniExperienceCancel struct {
}
type AlipayOpenMiniExperienceCreate struct {
}
type AlipayOpenMiniExperienceQuery struct {
}
type AlipayOpenMiniIndividualBusinessCertify struct {
}
type AlipayOpenMiniQrcodeBind struct {
}
type AlipayOpenMiniQrcodeUnbind struct {
}
type AlipayOpenMiniReleaststBatchquery struct {
}
type AlipayOpenMiniSafedomainCreate struct {
}
type AlipayOpenMiniSafedomainDelete struct {
}
type AlipayOpenMiniTemplatemessageUsertemplateApply struct {
}
type AlipayOpenMiniTemplateUsageQuery struct {
}
type AlipayOpenMiniTinyappExistQuery struct {
}
type AlipayOpenMiniVersionAuditApply struct {
}
type AlipayOpenMiniVersionAuditCancel struct {
}
type AlipayOpenMiniVersionAuditedCancel struct {
}
type AlipayOpenMiniVersionBuildQuery struct {
}
type AlipayOpenMiniVersionDelete struct {
}
type AlipayOpenMiniVersionDetailQuery struct {
}
type AlipayOpenMiniVersionGrayCancel struct {
}
type AlipayOpenMiniVersionGrayOnline struct {
}
type AlipayOpenMiniVersionListQuery struct {
}
type AlipayOpenMiniVersionOffline struct {
}
type AlipayOpenMiniVersionOnline struct {
}
type AlipayOpenMiniVersionRollback struct {
}
type AlipayOpenMiniVersionUpload struct {
}
type AlipayOpenNewgotoneCreate struct {
}
type AlipayOpenOperationBizfeeActivityApply struct {
}
type AlipayOpenOperationOpenbizmockBizQuery struct {
}
type AlipayOpenOperationSsffDeeQuery struct {
}
type AlipayOpenPageNewcontextTransfer struct {
}
type AlipayOpenPageOldcontextTransfer struct {
}
type AlipayOpenPublicAccountCreate struct {
}
type AlipayOpenPublicAccountDelete struct {
}
type AlipayOpenPublicAccountQuery struct {
}
type AlipayOpenPublicAccountReset struct {
}
type AlipayOpenPublicAdvertBatchquery struct {
}
type AlipayOpenPublicAdvertCreate struct {
}
type AlipayOpenPublicAdvertDelete struct {
}
type AlipayOpenPublicAdvertModify struct {
}
type AlipayOpenPublicArticlesummaryDataBatchquery struct {
}
type AlipayOpenPublicComptestCreate struct {
}
type AlipayOpenPublicContactFollowBatchquery struct {
}
type AlipayOpenPublicDefaultExtensionCreate struct {
}
type AlipayOpenPublicFollowBatchquery struct {
}
type AlipayOpenPublicGisQuery struct {
}
type AlipayOpenPublicGroupBatchquery struct {
}
type AlipayOpenPublicGroupCreate struct {
}
type AlipayOpenPublicGroupCrowdQuery struct {
}
type AlipayOpenPublicGroupDelete struct {
}
type AlipayOpenPublicGroupModify struct {
}
type AlipayOpenPublicInfoModify struct {
}
type AlipayOpenPublicInfoQuery struct {
}
type AlipayOpenPublicLabelCreate struct {
}
type AlipayOpenPublicLabelDelete struct {
}
type AlipayOpenPublicLabelModify struct {
}
type AlipayOpenPublicLabelQuery struct {
}
type AlipayOpenPublicLabelUserCreate struct {
}
type AlipayOpenPublicLabelUserDelete struct {
}
type AlipayOpenPublicLabelUserQuery struct {
}
type AlipayOpenPublicLifeAboardApply struct {
}
type AlipayOpenPublicLifeAccountCreate struct {
}
type AlipayOpenPublicLifeAgentcreateQuery struct {
}
type AlipayOpenPublicLifeAgentCreate struct {
}
type AlipayOpenPublicLifeCreate struct {
}
type AlipayOpenPublicLifeDebarkApply struct {
}
type AlipayOpenPublicLifeLabelBatchquery struct {
}
type AlipayOpenPublicLifeLabelCreate struct {
}
type AlipayOpenPublicLifeLabelDelete struct {
}
type AlipayOpenPublicLifeLabelModify struct {
}
type AlipayOpenPublicLifeModify struct {
}
type AlipayOpenPublicLifeMsgRecall struct {
}
type AlipayOpenPublicLifeMsgSend struct {
}
type AlipayOpenPublicMatchuserLabelCreate struct {
}
type AlipayOpenPublicMatchuserLabelDelete struct {
}
type AlipayOpenPublicMenuBatchquery struct {
}
type AlipayOpenPublicMenuCreate struct {
}
type AlipayOpenPublicMenuDataBatchquery struct {
}
type AlipayOpenPublicMenuDelete struct {
}
type AlipayOpenPublicMenuModify struct {
}
type AlipayOpenPublicMenuQuery struct {
}
type AlipayOpenPublicMessageContentCreate struct {
}
type AlipayOpenPublicMessageContentModify struct {
}
type AlipayOpenPublicMessageCustomSend struct {
}
type AlipayOpenPublicMessageGroupSend struct {
}
type AlipayOpenPublicMessageLabelSend struct {
}
type AlipayOpenPublicMessageQuery struct {
}
type AlipayOpenPublicMessageSingleSend struct {
}
type AlipayOpenPublicMessageTotalSend struct {
}
type AlipayOpenPublicMultimediaDownloadProxy struct {
}
type AlipayOpenPublicPartnerMenuOperate struct {
}
type AlipayOpenPublicPartnerMenuQuery struct {
}
type AlipayOpenPublicPartnerSubscribeSync struct {
}
type AlipayOpenPublicPayeeBindCreate struct {
}
type AlipayOpenPublicPayeeBindDelete struct {
}
type AlipayOpenPublicPersonalizedExtensionBatchquery struct {
}
type AlipayOpenPublicPersonalizedExtensionCreate struct {
}
type AlipayOpenPublicPersonalizedExtensionDelete struct {
}
type AlipayOpenPublicPersonalizedExtensionSet struct {
}
type AlipayOpenPublicPersonalizedMenuCreate struct {
}
type AlipayOpenPublicPersonalizedMenuDelete struct {
}
type AlipayOpenPublicQrcodeCreate struct {
}
type AlipayOpenPublicSettingCategoryQuery struct {
}
type AlipayOpenPublicShortlinkCreate struct {
}
type AlipayOpenPublicSinglearticleDataBatchquery struct {
}
type AlipayOpenPublicTemplateMessageAdd struct {
}
type AlipayOpenPublicTemplateMessageGet struct {
}
type AlipayOpenPublicTemplateMessageIndustryModify struct {
}
type AlipayOpenPublicThirdCustomerService struct {
}
type AlipayOpenPublicTopicBatchquery struct {
}
type AlipayOpenPublicTopicCreate struct {
}
type AlipayOpenPublicTopicDelete struct {
}
type AlipayOpenPublicTopicModify struct {
}
type AlipayOpenPublicUserDataBatchquery struct {
}
type AlipayOpenPublicUserFollowQuery struct {
}
type AlipayOpenPublicXwbtestabcdBatchquery struct {
}
type AlipayOpenServicemarketCommodityShopOffline struct {
}
type AlipayOpenServicemarketCommodityShopOnline struct {
}
type AlipayOpenServicemarketOrderAccept struct {
}
type AlipayOpenServicemarketOrderItemCancel struct {
}
type AlipayOpenServicemarketOrderItemComplete struct {
}
type AlipayOpenServicemarketOrderItemConfirm struct {
}
type AlipayOpenServicemarketOrderNotify struct {
}
type AlipayOpenServicemarketOrderQuery struct {
}
type AlipayOpenServicemarketOrderReject struct {
}
type AlipayOpenSmsgDataSet struct {
}
type AlipayOpenTestQuery struct {
}
type AlipayPassCodeAdd struct {
}
type AlipayPassCodeVerify struct {
}
type AlipayPassFileAdd struct {
}
type AlipayPassInstanceAdd struct {
}
type AlipayPassInstanceUpdate struct {
}
type AlipayPassSyncAdd struct {
}
type AlipayPassSyncUpdate struct {
}
type AlipayPassTemplateAdd struct {
}
type AlipayPassTemplateUpdate struct {
}
type AlipayPassTplAdd struct {
}
type AlipayPassTplContentAdd struct {
}
type AlipayPassTplContentUpdate struct {
}
type AlipayPassTplUpdate struct {
}
type AlipayPassVerifyQuery struct {
}
type AlipayPayCodecHschoolDecodeUse struct {
}
type AlipayPcreditHuabeiAuthAccumulationQuery struct {
}
type AlipayPcreditHuabeiAuthAgreementClose struct {
}
type AlipayPcreditHuabeiAuthAgreementQuery struct {
}
type AlipayPcreditHuabeiAuthBusinessConfirm struct {
}
type AlipayPcreditHuabeiAuthOrderQuery struct {
}
type AlipayPcreditHuabeiAuthRefundApply struct {
}
type AlipayPcreditHuabeiAuthSettleApply struct {
}
type AlipayPcreditHuabeiAuthSignApply struct {
}
type AlipayPcreditHuabeiDiscountBillQuery struct {
}
type AlipayPcreditHuabeiDiscountSolutionCreate struct {
}
type AlipayPcreditHuabeiDiscountSolutionModify struct {
}
type AlipayPcreditHuabeiDiscountSolutionOffline struct {
}
type AlipayPcreditHuabeiDiscountSolutionOnline struct {
}
type AlipayPcreditHuabeiPromoQuery struct {
}
type AlipayPcreditLoanApplyCreate struct {
}
type AlipayPcreditLoanRefundCreate struct {
}
type AlipayPlatformOpenidGet struct {
}
type AlipayPlatformUseridGet struct {
}
type AlipayPointBalanceGet struct {
}
type AlipayPointBudgetGet struct {
}
type AlipayPointOrderAdd struct {
}
type AlipayPointOrderGet struct {
}
type AlipayPromorulecenterRuleAnalyze struct {
}
type AlipaySecurityInfoAnalysis struct {
}
type AlipaySecurityProdAlipaySecurityProdTest struct {
}
type AlipaySecurityProdAmlriskQuery struct {
}
type AlipaySecurityProdCheckIqQuery struct {
}
type AlipaySecurityProdDdsBatchquery struct {
}
type AlipaySecurityProdDesQuery struct {
}
type AlipaySecurityProdDfasfdasFdfdsBatchquery struct {
}
type AlipaySecurityProdDfesfDefBatchquery struct {
}
type AlipaySecurityProdFacerepoAdd struct {
}
type AlipaySecurityProdFacerepoSearch struct {
}
type AlipaySecurityProdFingerprintApplyInitialize struct {
}
type AlipaySecurityProdFingerprintApply struct {
}
type AlipaySecurityProdFingerprintDelete struct {
}
type AlipaySecurityProdFingerprintDeviceVerify struct {
}
type AlipaySecurityProdFingerprintRiskcontrolQuery struct {
}
type AlipaySecurityProdFingerprintVerifyInitialize struct {
}
type AlipaySecurityProdFingerprintVerify struct {
}
type AlipaySecurityProdNopidBatchquery struct {
}
type AlipaySecurityProdSignatureFileUpload struct {
}
type AlipaySecurityProdSignatureTaskApply struct {
}
type AlipaySecurityProdSignatureTaskCancel struct {
}
type AlipaySecurityProdSignatureTaskQuery struct {
}
type AlipaySecurityProdXwbtestabcAbcQuery struct {
}
type AlipaySecurityProdXwbtestprodQuery struct {
}
type AlipaySecurityRiskContentDetect struct {
}
type AlipaySecurityRiskCustomerriskQuery struct {
}
type AlipaySecurityRiskCustomerriskSend struct {
}
type AlipaySecurityRiskDetect struct {
}
type AlipaySecurityRiskHideDeviceidQuery struct {
}
type AlipaySecurityRiskRainscoreQuery struct {
}
type AlipaySecuritySssssssQuery struct {
}
type AlipaySecurityTesttestQuery struct {
}
type AlipaySystemOauthToken struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}
type ResponseAlipaySystemOauthToken struct {
	AlipaySystemOauthTokenResponse *AlipaySystemOauthTokenResponse `json:"alipay_system_oauth_token_response"`
	ErrorResponse                  *ErrorResponse                  `json:"error_response"`
	Sign                           string                          `json:"sign"`
}
type AlipaySystemOauthTokenResponse struct {
	*ErrorResponse
	UserId       string `json:"user_id"`       //支付宝用户的唯一userId
	AccessToken  string `json:"access_token"`  //访问令牌,通过该令牌调用需要授权类接口
	ExpiresIn    int64  `json:"expires_in"`    //访问令牌的有效时间，单位是秒。
	ReExpiresIn  int64  `json:"re_expires_in"` //刷新令牌。通过该令牌可以刷新access_token
	AlipayUserId string `json:"alipay_user_id"`
	RefreshToken string `json:"refresh_token"` //刷新令牌的有效时间，单位是秒
}

type Response struct {
	ErrorResponse *ErrorResponse `json:"error_response"`
	Sign          string         `json:"sign"`
}
type ErrorResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func (e *ErrorResponse) Error() error {
	if e != nil && e.Code != "10000" && e.Code != "" {
		return errors.New(e.SubCode + ":" + e.SubMsg)
	}
	return nil
}

type AlipayTradeAdvanceConsult struct {
}
type AlipayTradeAppPay struct {
	AlipayTradeAppPayResponse AlipayTradeWapPayResponse `json:"alipay_trade_app_pay_response"`
	Sign                      string                    `json:"sign"`
}
type AlipayTradeWapPay struct {
	AlipayTradeWapPayResponse AlipayTradeWapPayResponse `json:"alipay_trade_wap_pay_response"`
	Sign                      string                    `json:"sign"`
}
type AlipayTradeWapPayResponse struct {
	*ErrorResponse
	OutTradeNo      string  `json:"out_trade_no"`      //64 	商户网站唯一订单号 	70501111111S001111119
	TradeNo         string  `json:"trade_no"`          //64 	该交易在支付宝系统中的交易流水号。最长64位。 	2014112400001000340011111118
	TradeAmount     float64 `json:"trade_amount"`      //9 	该笔订单的资金总额，单位为RMB-Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。 	9.00
	SellerId        string  `json:"seller_id"`         //16 	收款支付宝账号对应的支付宝唯一用户号。 以2088开头的纯16位数字 	2088111111116894
	MerchantOrderNo string  `json:"merchant_order_no"` //32 	商户原始订单号，最大长度限制32位 	20161008001
}
type AlipayTradeCreate struct {
	AlipayTradeCreateResponse AlipayTradeCreateResponse
	Sign                      string `json:"sign"`
}
type AlipayTradeCreateResponse struct {
	*ErrorResponse
	OutTradeNo string `json:"out_trade_no"` //64 	商户网站唯一订单号 	70501111111S001111119
	TradeNo    string `json:"trade_no"`     //64 	该交易在支付宝系统中的交易流水号。最长64位。 	2014112400001000340011111118
}
type AlipayTradeCancel struct {
	AlipayTradeCancelResponse AlipayTradeCancelResponse `json:"alipay_trade_cancel_response"`
	Sign                      string                    `json:"sign"`
}
type AlipayTradeCancelResponse struct {
	*ErrorResponse
	TradeNo             string `json:"trade_no"`             //64 	支付宝交易号; 当发生交易关闭或交易退款时返回； 	2013112011001004330000121536
	OutTradeNo          string `json:"out_trade_no"`         //64 	商户订单号 	6823789339978248
	RetryFlag           string `json:"retry_flag"`           //1 	是否需要重试 	N
	Action              string `json:"action"`               //10 	本次撤销触发的交易动作,接口调用成功且交易存在时返回。可能的返回值： close：交易未支付，触发关闭交易动作，无退款； refund：交易已支付，触发交易退款动作； 未返回：未查询到交易，或接口调用失败； 	Close
	GmtRefundPay        string `json:"gmt_refund_pay"`       //32 	当撤销产生了退款时，返回退款时间； 默认不返回该信息，需与支付宝约定后配置返回
	RefundSettlement_id string `json:"refund_settlement_id"` //64 	当撤销产生了退款时，返回的退款清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息； 	2018101610032004620239146945
}
type AlipayTradeClose struct {
	AlipayTradeCloseResponse AlipayTradeCloseResponse `json:"alipay_trade_close_response"`
	Sign                     string                   `json:"sign"`
}
type AlipayTradeCloseResponse struct {
	*ErrorResponse
	TradeNo    string `json:"trade_no"`     //64 	支付宝交易号; 当发生交易关闭或交易退款时返回； 	2013112011001004330000121536
	OutTradeNo string `json:"out_trade_no"` //64 	商户订单号 	6823789339978248
}
type AlipayTradeCustomsDeclare struct {
}
type AlipayTradeCustomsQuery struct {
}
type AlipayTradeFastpayRefundQuery struct {
	AlipayTradeFastpayRefundQueryResponse AlipayTradeFastpayRefundQueryResponse `json:"alipay_trade_fastpay_refund_query_response"`
	Sign                                  string                                `json:"sign"`
}

type AlipayTradeFastpayRefundQueryResponse struct {
	*ErrorResponse
	TradeNo                      string               `json:"trade_no"`       //必选 64 	商户网站唯一订单号
	OutTradeNo                   string               `json:"out_trade_no"`   //必选 64 	商户网站唯一订单号
	OutRequestNo                 string               `json:"out_request_no"` //64 本笔退款对应的退款请求号
	RefundReason                 string               `json:"refund_reason"`  //发起退款时，传入的退款原因 	用户退款请求
	TotalAmount                  float64              `json:"total_amount"`   //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	RefundAmount                 float64              `json:"refund_amount"`  //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	RefundRoyaltys               *RefundRoyaltyResult `json:"refund_royaltys"`
	GmtRefundPay                 string               `json:"gmt_refund_pay"`                  //32 	退款时间； 默认不返回该信息，需与支付宝约定后配置返回； 	2014-11-27 15:45:57
	RefundDetailItemList         *TradeFundBill       `json:"refund_detail_item_list"`         //本次退款使用的资金渠道； 默认不返回该信息，需与支付宝约定后配置，或者入参的query_options中指定时才返回该字段信息。
	SendBackFee                  string               `json:"send_back_fee"`                   //11 	本次商户实际退回金额； 默认不返回该信息，需与支付宝约定后配置返回； 	88
	RefundSettlementId           string               `json:"refund_settlement_id"`            //退款清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息； 	2018101610032004620239146945
	PresentRefundBuyerAmount     string               `json:"present_refund_buyer_amount"`     //本次退款金额中买家退款金额
	PresentRefundDiscountAmount  string               `json:"present_refund_discount_amount"`  //本次退款金额中平台优惠退款金额
	PresentRefundMdiscountAmount string               `json:"present_refund_mdiscount_amount"` //本次退款金额中商家优惠退款金额
}
type RefundRoyaltyResult struct {
	RefundAmount  float64 `json:"refund_amount"`   //9 	退分账金额
	RoyaltyType   string  `json:"royalty_type"`    //32 	分账类型. 普通分账为：transfer; 补差为：replenish; 为空默认为分账transfer; 	Transfer
	ResultCode    string  `json:"result_code"`     //32 	退分账结果码 	SUCCESS
	TransOut      string  `json:"trans_out"`       //28 	转出人支付宝账号对应用户ID 	2088102210397302
	TransOutEmail string  `json:"trans_out_email"` //64 	转出人支付宝账号 	Alipay-test03@alipay.com
	TransIn       string  `json:"trans_in"`        //28 	转入人支付宝账号对应用户ID 	2088102210397302
	TransInEmail  string  `json:"trans_in_email"`  //64 	转入人支付宝账号 	zen_gwen@hotmail.com
}

type AlipayTradeOrderinfoSync struct {
}
type AlipayTradeOrderSettle struct {
}
type AlipayTradePageRefund struct {
	AlipayTradePageRefundResponse AlipayTradePageRefundResponse `json:"alipay_trade_page_refund_response"`
	Sign                          string                        `json:"sign"`
}
type AlipayTradePageRefundResponse struct {
	*ErrorResponse
	TradeNo         string  `json:"trade_no"`          //必选 64 	商户网站唯一订单号
	OutTradeNo      string  `json:"out_trade_no"`      //必选 64 	商户网站唯一订单号
	SellerId        string  `json:"seller_id"`         //28 	收款支付宝账号对应的支付宝唯一用户号。 以2088开头的纯16位数字 	2088111111116894
	TotalAmount     float64 `json:"total_amount"`      //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	MerchantOrderNo string  `json:"merchant_order_no"` //32 	商户原始订单号，最大长度限制32位 	20161008001
}
type AlipayTradePay struct {
	AlipayTradePayResponse AlipayTradePayResponse
	Sign                   string `json:"sign"`
}
type AlipayTradePayResponse struct {
	*ErrorResponse
	TradeNo             string         `json:"trade_no"`          //必选 64 	商户网站唯一订单号
	OutTradeNo          string         `json:"out_trade_no"`      //必选 64 	商户网站唯一订单号
	OutRequestNo        string         `json:"out_request_no"`    //64 本笔退款对应的退款请求号
	BuyerLogonId        string         `json:"buyer_logon_id"`    //必填 	100 	买家支付宝账号
	RefundReason        string         `json:"refund_reason"`     //发起退款时，传入的退款原因 	用户退款请求
	SettleAmount        string         `json:"settle_amount"`     //发起退款时，传入的退款原因 	用户退款请求
	PayCurrency         string         `json:"pay_currency"`      //发起退款时，传入的退款原因 	用户退款请求
	PayAmount           string         `json:"pay_amount"`        //发起退款时，传入的退款原因 	用户退款请求
	SettleTransRate     string         `json:"settle_trans_rate"` //发起退款时，传入的退款原因 	用户退款请求
	TransPayRate        string         `json:"trans_pay_rate"`    //发起退款时，传入的退款原因 	用户退款请求
	TotalAmount         float64        `json:"total_amount"`      //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	TransCurrency       string         `json:"trans_currency"`    //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	SettleCurrency      string         `json:"settle_currency"`   //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	ReceiptAmount       string         `json:"receipt_amount"`    //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	BuyerPayAmount      float64        `json:"buyer_pay_amount"`  //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	PointAmount         float64        `json:"point_amount"`      //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	InvoiceAmount       float64        `json:"invoice_amount"`    //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	GmtPayment          float64        `json:"gmt_payment"`       //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	FundBillList        *TradeFundBill `json:"fund_bill_list"`    //本次退款使用的资金渠道； 默认不返回该信息，需与支付宝约定后配置，或者入参的query_options中指定时才返回该字段信息。
	CardBalance         float64        `json:"card_balance"`      //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	StoreName           string         `json:"store_name"`
	BuyerUserId         string         `json:"buyer_user_id"`
	DiscountGoodsDetail string         `json:"discount_goods_detail"`
	VoucherDetailList   *VoucherDetail `json:"voucher_detail_list"`
	AdvanceAmount       string         `json:"advance_amount"`
	AuthTradePayMode    string         `json:"auth_trade_pay_mode"` //32 	退款时间； 默认不返回该信息，需与支付宝约定后配置返回； 	2014-11-27 15:45:57
	ChargeAmount        string         `json:"charge_amount"`       //11 	本次商户实际退回金额； 默认不返回该信息，需与支付宝约定后配置返回； 	88
	ChargeFlags         string         `json:"charge_flags"`        //退款清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息； 	2018101610032004620239146945
	SettlementId        string         `json:"settlement_id"`       //本次退款金额中买家退款金额
	BusinessParams      string         `json:"business_params"`     //本次退款金额中平台优惠退款金额
	BuyerUserType       string         `json:"buyer_user_type"`     //本次退款金额中商家优惠退款金额
	MdiscountAmount     string         `json:"mdiscount_amount"`    //本次退款金额中商家优惠退款金额
	DiscountAmount      string         `json:"discount_amount"`     //本次退款金额中商家优惠退款金额
	BuyerUserName       string         `json:"buyer_user_name"`     //选填 	128 	买家名称； 买家为个人用户时为买家姓名，买家为企业用户时为企业名称；默认不返回该信息，需与支付宝约定后配置返回； 	菜鸟网络有限公司
}
type AlipayTradePrecreateResponse struct {
	*ErrorResponse
	OutTradeNo string `json:"out_trade_no"` //64 	商户的订单号 	6823789339978248
	ArCode     string `json:"qr_code"`      //1024 	当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码 	https://qr.alipay.com/bavh4wjlxf12tper3a
}
type AlipayTradePrecreate struct {
	AlipayTradePrecreateResponse *AlipayTradePrecreateResponse `json:"alipay_trade_precreate_response"`
	Sign                         string                        `json:"sign"`
}
type AlipayTradePagePay struct {
	AlipayTradePagePayResponse AlipayTradePagePayResponse `json:"alipay_trade_page_pay_response"`
	Sign                       string                     `json:"sign"`
}
type AlipayTradePagePayResponse struct {
	*ErrorResponse
	TradeNo         string  `json:"trade_no"`          //必选 64 	商户网站唯一订单号
	OutTradeNo      string  `json:"out_trade_no"`      //必选 64 	商户网站唯一订单号
	SellerId        string  `json:"seller_id"`         //28 	收款支付宝账号对应的支付宝唯一用户号。 以2088开头的纯16位数字 	2088111111116894
	TotalAmount     float64 `json:"total_amount"`      //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	MerchantOrderNo string  `json:"merchant_order_no"` //32 	商户原始订单号，最大长度限制32位 	20161008001
}

type AlipayTradeQuery struct {
	AlipayTradeQueryResponse AlipayTradeQueryResponse `json:"alipay_trade_query_response"`
	Sign                     string                   `json:"sign"`
}
type AlipayTradeQueryResponse struct {
	*ErrorResponse
	TradeNo       string  `json:"trade_no"`        //必选 64 	商户网站唯一订单号
	OutTradeNo    string  `json:"out_trade_no"`    //必选 64 	商户网站唯一订单号
	BuyerLogonId  string  `json:"buyer_logon_id"`  //必填 	100 	买家支付宝账号 	159****5620
	TradeStatus   string  `json:"trade_status"`    //必填 	32 	交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款） 	TRADE_CLOSED
	TotalAmount   float64 `json:"total_amount"`    //必填 	11 	交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount 	88.88
	StoreName     string  `json:"store_name"`      //512 	请求交易支付中的商户店铺的名称 	证大五道口店
	BuyerUserId   string  `json:"buyer_user_id"`   //买家在支付宝的用户id 	2088101117955611
	BuyerUserType string  `json:"buyer_user_type"` //18 	买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。 	PRIVATE
}
type AlipayTradeRefund struct {
	AlipayTradeRefundResponse AlipayTradeRefundResponse `json:"alipay_trade_refund_response"`
	Sign                      string                    `json:"sign"`
}
type AlipayTradeRefundResponse struct {
	*ErrorResponse
	TradeNo                      string             `json:"trade_no"`                        //必选 64 	商户网站唯一订单号
	OutTradeNo                   string             `json:"out_trade_no"`                    //必选 64 	商户网站唯一订单号
	BuyerLogonId                 string             `json:"buyer_logon_id"`                  //必填 	100 	买家支付宝账号 	159****5620
	FundChange                   string             `json:"fund_change"`                     //1 	本次退款是否发生了资金变化 	Y
	RefundFee                    float64            `json:"refund_fee"`                      //11 	退款总金额 	88.88
	RefundCurrency               float64            `json:"refund_currency"`                 //退款币种信息
	GmtRefundPay                 string             `json:"gmt_refund_pay"`                  //退款支付时间
	RefundDetailItemList         *TradeFundBill     `json:"refund_detail_item_list"`         //退款使用的资金渠道。 只有在签约中指定需要返回资金明细，或者入参的query_options中指定时才返回该字段信息。
	StoreName                    string             `json:"store_name"`                      //512 	请求交易支付中的商户店铺的名称 	证大五道口店
	BuyerUserId                  string             `json:"buyer_user_id"`                   //买家在支付宝的用户id 	2088101117955611
	RefundPresetPaytoolList      *PresetPayToolInfo `json:"refund_preset_paytool_list"`      //退回的前置资产列表
	RefundSettlementId           string             `json:"refund_settlement_id"`            //退款清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息； 	2018101610032004620239146945
	PresentRefundBuyerAmount     string             `json:"present_refund_buyer_amount"`     //本次退款金额中买家退款金额
	PresentRefundDiscountAmount  string             `json:"present_refund_discount_amount"`  //本次退款金额中平台优惠退款金额
	PresentRefundMdiscountAmount string             `json:"present_refund_mdiscount_amount"` //本次退款金额中商家优惠退款金额

}
type PresetPayToolInfo struct {
	Amount         []float64 `json:"amount"`           //前置资产金额
	AssertTypeCode string    `json:"assert_type_code"` //前置资产类型编码，和收单支付传入的preset_pay_tool里面的类型编码保持一致。 	盒马礼品卡:HEMA
	// ；抓猫猫红包:T_CAT_COUPON
}
type AlipayTradeRoyaltyRelationBatchquery struct {
}
type AlipayTradeRoyaltyRelationBind struct {
}
type AlipayTradeRoyaltyRelationUnbind struct {
}
type AlipayTradeVendorpayDevicedataUpload struct {
}
type AlipayTransferThirdpartyBillCreate struct {
}
type AlipayTrustUserAuthSend struct {
}
type AlipayTrustUserReportGet struct {
}
type AlipayTrustUserRiskidentifyGet struct {
}
type AlipayTrustUserScoreGet struct {
}
type AlipayTrustUserTokenGet struct {
}
type AlipayUserAccountFreezeGet struct {
}
type AlipayUserAccountGet struct {
}
type AlipayUserAccountSearch struct {
}
type AlipayUserAccountUseridBatchquery struct {
}
type AlipayUserAddressQuery struct {
}
type AlipayUserAgreementExecutionplanModify struct {
}
type AlipayUserAgreementPageSign struct {
}
type AlipayUserAgreementQuery struct {
}
type AlipayUserAgreementTransfer struct {
}
type AlipayUserAgreementUnsign struct {
}
type AlipayUserAlipaypointSend struct {
}
type AlipayUserAuthZhimaorgIdentityApply struct {
}
type AlipayUserCertifyOpenCertify struct {
}
type AlipayUserCertifyOpenInitialize struct {
}
type AlipayUserCertifyOpenQuery struct {
}
type AlipayUserContractGet struct {
}
type AlipayUserFinanceinfoShare struct {
}
type AlipayUserGet struct {
}
type AlipayUserInfoAuth struct {
	Scopes []string `json:"scopes"` //接口权限值，目前只支持auth_user和auth_base两个值
	/*
		Auth_base：以auth_base为scope发起的网页授权，是用来获取进入页面的用户的userId的，并且是静默授权并自动跳转到回调页的。用户感知的就是直接进入了回调页（通常是业务页面）。
		Auth_user：以auth_user为scope发起的网页授权，是用来获取用户的基本信息的（比如头像、昵称等）。但这种授权需要用户手动同意，用户同意后，就可在授权后获取到该用户的基本信息
	*/
	State string `json:"state"` //商户自定义参数，用户授权后，重定向到redirect_uri时会原样回传给商户。 为防止CSRF攻击，建议开发者请求授权时传入state参数，该参数要做到既不可预测，又可以证明客户端和当前第三方网站的登录认证状态存在关联。 只允许base64字符（长度小于等于100）
}
type AlipayUserInfoShareResponse struct {
	*ErrorResponse
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	Gender   string `json:"gender"`
	NickName string `json:"nick_name"`
	Province string `json:"province"`
	UserId   string `json:"user_id"`
}
type AlipayUserInfoShare struct {
	AlipayUserInfoShareResponse *AlipayUserInfoShareResponse `json:"alipay_user_info_share_response"`
	Sign                        string                       `json:"sign"`
} //alipay.user.info.share  付宝会员授权信息查询接口
type AlipayUserPassInstancebatchAdd struct {
}
type AlipayUserPassTemplateCreate struct {
}
type AlipayUserPassTemplateModify struct {
}
type AlipayUserPassTemplateQuery struct {
}
type AlipayUserTest struct {
}
type AlipayUserTradeSearch struct {
}
type AlipayUserTwostageCommonUse struct {
}
type AlipayUserUserinfoShare struct {
}
type AlipayZdataassetsEasyservice struct {
}
type AlipayZdataassetsFcdatalabZdatamergetask struct {
}
type AlipayZdataassetsMetadata struct {
}
type AlipayZdatafrontCommonQuery struct {
}
type AlipayZdatafrontDatatransferedFileupload struct {
}
type AlipayZdatafrontDatatransferedSend struct {
}
type AlipayZdataserviceUnidataQuery struct {
}
type AlipayZmscoreZrankGet struct {
}
type AmapMapMapserviceTeseBatchquery struct {
}
type AntMerchantExpandContractFacetofaceQuery struct {
}
type AntMerchantExpandContractFacetofaceSign struct {
}
type AntMerchantExpandEnterpriseApply struct {
}
type AntMerchantExpandImageUpload struct {
}
type AntMerchantExpandItemOpenBatchquery struct {
}
type AntMerchantExpandItemOpenCreate struct {
}
type AntMerchantExpandItemOpenDelete struct {
}
type AntMerchantExpandItemOpenModify struct {
}
type AntMerchantExpandItemOpenQuery struct {
}
type AntMerchantExpandMapplyorderQuery struct {
}
type AntMerchantExpandMerchantStorelistQuery struct {
}
type AntMerchantExpandPersonalApply struct {
}
type AntOcrDriverlicenseIdentify struct {
}
type AntOcrGeneralIdentify struct {
}
type AntOcrIdcardIdentify struct {
}
type AntOcrVehiclelicenseIdentify struct {
}
type AntOcrVehicleplateIdentify struct {
}
type AnttechBlockchainFinanceFileUpload struct {
}
type KoubeiAdvertDeliveryDiscountAuthwebBatchquery struct {
}
type KoubeiAdvertDeliveryDiscountGet struct {
}
type KoubeiAdvertDeliveryDiscountQuery struct {
}
type KoubeiAdvertDeliveryDiscountSend struct {
}
type KoubeiAdvertDeliveryDiscountWebBatchquery struct {
}
type KoubeiAdvertDeliveryItemApply struct {
}
type KoubeiCateringCommodityOrderBuy struct {
}
type KoubeiCateringCrowdgroupConditionQuery struct {
}
type KoubeiCateringCrowdgroupConditionSet struct {
}
type KoubeiCateringDishCommgroupQuery struct {
}
type KoubeiCateringDishCommgroupSync struct {
}
type KoubeiCateringDishCommruleQuery struct {
}
type KoubeiCateringDishCommruleSync struct {
}
type KoubeiCateringDishCookcatetopSync struct {
}
type KoubeiCateringDishSpecgroupQuery struct {
}
type KoubeiCateringDishSpecgroupSync struct {
}
type KoubeiCateringDishVirtualdishQuery struct {
}
type KoubeiCateringDishVirtualdishSync struct {
}
type KoubeiCateringTablecodeQuery struct {
}
type KoubeiCateringTablelistQuery struct {
}
type KoubeiCraftsmanDataProviderBatchquery struct {
}
type KoubeiCraftsmanDataProviderCreate struct {
}
type KoubeiCraftsmanDataProviderModify struct {
}
type KoubeiCraftsmanDataWorkBatchquery struct {
}
type KoubeiCraftsmanDataWorkCreate struct {
}
type KoubeiCraftsmanDataWorkDelete struct {
}
type KoubeiCraftsmanDataWorkModify struct {
}
type KoubeiItemCategoryChildrenBatchquery struct {
}
type KoubeiItemCreate struct {
}
type KoubeiItemExtitemBatchquery struct {
}
type KoubeiItemExtitemBrandQuery struct {
}
type KoubeiItemExtitemCategoryQuery struct {
}
type KoubeiItemExtitemCreate struct {
}
type KoubeiItemExtitemExistedQuery struct {
}
type KoubeiItemExtitemQuery struct {
}
type KoubeiItemExtitemUpdate struct {
}
type KoubeiItemModify struct {
}
type KoubeiItemState struct {
}
type KoubeiMarketingCampaignActivityBatchquery struct {
}
type KoubeiMarketingCampaignActivityCreate struct {
}
type KoubeiMarketingCampaignActivityModify struct {
}
type KoubeiMarketingCampaignActivityOffline struct {
}
type KoubeiMarketingCampaignActivityQuery struct {
}
type KoubeiMarketingCampaignAssetDetailQuery struct {
}
type KoubeiMarketingCampaignCrowdBatchquery struct {
}
type KoubeiMarketingCampaignCrowdCount struct {
}
type KoubeiMarketingCampaignCrowdCreate struct {
}
type KoubeiMarketingCampaignCrowdDelete struct {
}
type KoubeiMarketingCampaignCrowdDetailQuery struct {
}
type KoubeiMarketingCampaignCrowdModify struct {
}
type KoubeiMarketingCampaignDetailInfoQuery struct {
}
type KoubeiMarketingCampaignIntelligentPromoBatchquery struct {
}
type KoubeiMarketingCampaignIntelligentPromoConsult struct {
}
type KoubeiMarketingCampaignIntelligentPromoCreate struct {
}
type KoubeiMarketingCampaignIntelligentPromoDelete struct {
}
type KoubeiMarketingCampaignIntelligentPromoModify struct {
}
type KoubeiMarketingCampaignIntelligentPromoQuery struct {
}
type KoubeiMarketingCampaignIntelligentShopConsult struct {
}
type KoubeiMarketingCampaignIntelligentTemplateConsult struct {
}
type KoubeiMarketingCampaignItemMerchantactivityBatchquery struct {
}
type KoubeiMarketingCampaignItemMerchantactivityClose struct {
}
type KoubeiMarketingCampaignItemMerchantactivityCreate struct {
}
type KoubeiMarketingCampaignItemMerchantactivityModify struct {
}
type KoubeiMarketingCampaignItemMerchantactivityQuery struct {
}
type KoubeiMarketingCampaignRecruitApplyQuery struct {
}
type KoubeiMarketingCampaignRecruitShopQuery struct {
}
type KoubeiMarketingCampaignTagsQuery struct {
}
type KoubeiMarketingCampaignUserAssetQuery struct {
}
type KoubeiMarketingDataActivityBillDownload struct {
}
type KoubeiMarketingDataActivityReportQuery struct {
}
type KoubeiMarketingDataAlisisReportBatchquery struct {
}
type KoubeiMarketingDataAlisisReportQuery struct {
}
type KoubeiMarketingDataBizadviserMemberprofileQuery struct {
}
type KoubeiMarketingDataBizadviserMyddsreportQuery struct {
}
type KoubeiMarketingDataBizadviserMyreportQuery struct {
}
type KoubeiMarketingDataCustomreportBatchquery struct {
}
type KoubeiMarketingDataCustomreportDelete struct {
}
type KoubeiMarketingDataCustomreportDetailQuery struct {
}
type KoubeiMarketingDataCustomreportQuery struct {
}
type KoubeiMarketingDataCustomreportSave struct {
}
type KoubeiMarketingDataDishdiagnoseBatchquery struct {
}
type KoubeiMarketingDataDishdiagnosetypeBatchquery struct {
}
type KoubeiMarketingDataEnterpriseStaffinfoUpload struct {
}
type KoubeiMarketingDataIndicatorQuery struct {
}
type KoubeiMarketingDataIntelligentEffectQuery struct {
}
type KoubeiMarketingDataIntelligentIndicatorQuery struct {
}
type KoubeiMarketingDataIsvShopQuery struct {
}
type KoubeiMarketingDataMallShopitemsQuery struct {
}
type KoubeiMarketingDataMemberReportQuery struct {
}
type KoubeiMarketingDataMessageDeliver struct {
}
type KoubeiMarketingDataNearmallQuery struct {
}
type KoubeiMarketingDataRetailDmQuery struct {
}
type KoubeiMarketingDataSmartactivityConfig struct {
}
type KoubeiMarketingDataSmartactivityForecast struct {
}
type KoubeiMarketingDataSmartmanagementDiagnose struct {
}
type KoubeiMarketingDataTradeHabbitQuery struct {
}
type KoubeiMarketingToolIsvMerchantQuery struct {
}
type KoubeiMarketingToolPointsQuery struct {
}
type KoubeiMarketingToolPointsUpdate struct {
}
type KoubeiMarketingToolPrizesendAuth struct {
}
type KoubeiMemberBrandownerNameQuery struct {
}
type KoubeiMemberDataDesdBatchquery struct {
}
type KoubeiMemberDataIsvCreate struct {
}
type KoubeiMemberDataOauthQuery struct {
}
type KoubeiMemberRetailerQuery struct {
}
type KoubeiMerchantKbdeviceDevicesBatchquery struct {
}
type KoubeiMerchantKbdeviceDispenserQuery struct {
}
type KoubeiQualityTestCloudacptActivityQuery struct {
}
type KoubeiQualityTestCloudacptBatchQuery struct {
}
type KoubeiQualityTestCloudacptCheckresultSubmit struct {
}
type KoubeiQualityTestCloudacptItemQuery struct {
}
type KoubeiRetailShopitemBatchquery struct {
}
type KoubeiRetailShopitemModify struct {
}
type KoubeiRetailShopitemUpload struct {
}
type KoubeiSalesLeadsSaleleadsCreate struct {
}
type KoubeiSalesLeadsShopleadsCreate struct {
}
type KoubeiTradeItemBuy struct {
}
type KoubeiTradeItemorderBuy struct {
}
type KoubeiTradeItemorderQuery struct {
}
type KoubeiTradeItemorderRefund struct {
}
type KoubeiTradeKbdeliveryDeliveryApply struct {
}
type KoubeiTradeKbdeliveryDeliveryCancel struct {
}
type KoubeiTradeOrderAggregateConsult struct {
}
type KoubeiTradeOrderConsult struct {
}
type KoubeiTradeOrderEnterpriseQuery struct {
}
type KoubeiTradeOrderEnterpriseSettle struct {
}
type KoubeiTradeOrderPrecreate struct {
}
type KoubeiTradeOrderQuery struct {
}
type KoubeiTradeTicketSendClose struct {
}
type KoubeiTradeTicketTicketcodeCancel struct {
}
type KoubeiTradeTicketTicketcodeCheckavailable struct {
}
type KoubeiTradeTicketTicketcodeDelay struct {
}
type KoubeiTradeTicketTicketcodeQuery struct {
}
type KoubeiTradeTicketTicketcodeSend struct {
}
type KoubeiTradeTicketTicketcodeSync struct {
}
type KoubeiTradeTicketTicketcodeUse struct {
}
type MonitorHeartbeatSyn struct {
}
type MybankCreditLoanapplyDataUpload struct {
}
type MybankFinanceYulibaoAccountQuery struct {
}
type MybankFinanceYulibaoCapitalPurchase struct {
}
type MybankFinanceYulibaoCapitalRansom struct {
}
type MybankFinanceYulibaoPriceQuery struct {
}
type MybankFinanceYulibaoTransHistoryQuery struct {
}
type MybankPaymentTradeFinancingOrderClose struct {
}
type MybankPaymentTradeFinancingOrderCreate struct {
}
type MybankPaymentTradeFinancingOrderRefund struct {
}
type MybankPaymentTradeOrderCreate struct {
}
type MybankPaymentTradeQrcodeCreate struct {
}
type MybankPaymentTradeQrcodeDelete struct {
}
type SsdataDataserviceRiskAlixiaohaoQuery struct {
}
type SsdataDataserviceRiskAntifraudlistQuery struct {
}
type SsdataDataserviceRiskAntifraudscoreQuery struct {
}
type ZhimaAuthInfoAuthquery struct {
}
type ZhimaCreditAntifraudRiskList struct {
}
type ZhimaCreditAntifraudScoreGet struct {
}
type ZhimaCreditAntifraudVerify struct {
}
type ZhimaCreditContractBorrowCancel struct {
}
type ZhimaCreditContractBorrowCreate struct {
}
type ZhimaCreditContractBorrowDelay struct {
}
type ZhimaCreditContractBorrowInitialize struct {
}
type ZhimaCreditContractBorrowQuery struct {
}
type ZhimaCreditContractBorrowReturn struct {
}
type ZhimaCreditContractPrincipalQuery struct {
}
type ZhimaCreditEpInfoGet struct {
}
type ZhimaCreditEpLawsuitDetailGet struct {
}
type ZhimaCreditEpLawsuitRecordGet struct {
}
type ZhimaCreditEpSceneAgreementCancel struct {
}
type ZhimaCreditEpSceneAgreementUse struct {
}
type ZhimaCreditEpSceneFulfillmentlistSync struct {
}
type ZhimaCreditEpSceneFulfillmentSync struct {
}
type ZhimaCreditEpSceneRatingApply struct {
}
type ZhimaCreditEpSceneRatingInitialize struct {
}
type ZhimaCreditEpSceneRatingQuery struct {
}
type ZhimaCreditEpScoreGet struct {
}
type ZhimaCreditOrderRepaymentApply struct {
}
type ZhimaCreditPeLawsuitDetailQuery struct {
}
type ZhimaCreditPeLawsuitRecordGet struct {
}
type ZhimaCreditScoreBriefGet struct {
}
type ZhimaCreditScoreGet struct {
}
type ZhimaCreditWatchlistBriefGet struct {
}
type ZhimaCreditWatchlistiiGet struct {
}
type ZhimaCustomerAuthMutualviewApply struct {
}
type ZhimaCustomerCertificationCertify struct {
}
type ZhimaCustomerCertificationInitialize struct {
}
type ZhimaCustomerCertificationMaterialCertify struct {
}
type ZhimaCustomerCertificationQuery struct {
}
type ZhimaCustomerContractInitialize struct {
}
type ZhimaCustomerEpCertificationCertify struct {
}
type ZhimaCustomerEpCertificationInitialize struct {
}
type ZhimaCustomerEpCertificationQuery struct {
}
type ZhimaDataBatchFeedback struct {
}
type ZhimaDataFeedbackurlQuery struct {
}
type ZhimaDataStateDataSync struct {
}
type ZhimaMerchantBorrowEntityUpload struct {
}
type ZhimaMerchantCloseloopDataUpload struct {
}
type ZhimaMerchantContractCommonCancel struct {
}
type ZhimaMerchantContractCommonConfirm struct {
}
type ZhimaMerchantContractCommonQuery struct {
}
type ZhimaMerchantContractOfferModify struct {
}
type ZhimaMerchantContractOfferQuery struct {
}
type ZhimaMerchantContractOnofferQuery struct {
}
type ZhimaMerchantContractPageQuery struct {
}
type ZhimaMerchantContractQuickCreate struct {
}
type ZhimaMerchantCreditserviceDetailCreate struct {
}
type ZhimaMerchantCreditserviceDetailModify struct {
}
type ZhimaMerchantCreditserviceDetailQuery struct {
}
type ZhimaMerchantDataUploadInitialize struct {
}
type ZhimaMerchantLogoImageUpload struct {
}
type ZhimaMerchantOrderRentCancel struct {
}
type ZhimaMerchantOrderRentComplete struct {
}
type ZhimaMerchantOrderRentCreate struct {
}
type ZhimaMerchantOrderRentModify struct {
}
type ZhimaMerchantOrderRentQuery struct {
}
type ZhimaMerchantSingleDataUpload struct {
}
type ZhimaMerchantTestPractice struct {
}
type ZhimaOpenAppDesSend struct {
}
type ZhimaOpenAppKeyanLqlCreate struct {
}
type ZhimaOpenQerqQerqQuery struct {
}
type ZolozAuthenticationCustomerFacemanageCreate struct {
}
type ZolozAuthenticationCustomerFacemanageDelete struct {
}
type ZolozAuthenticationCustomerFtokenQuery struct {
}
type ZolozAuthenticationCustomerSmilepayInitialize struct {
}
type ZolozAuthenticationSmilepayInitialize struct {
}
type ZolozIdentificationCustomerCertifyzhubQuery struct {
}
type ZolozIdentificationUserWebInitialize struct {
}
type ZolozIdentificationUserWebQuery struct {
}

type TradeFundBill struct {
	FundChannel string  `json:"fund_channel"` //32 	交易使用的资金渠道，详见 支付渠道列表 	ALIPAYACCOUNT
	BankCode    string  `json:"bank_code"`    //10 	银行卡支付时的银行代码 	CEB
	Amount      float64 `json:"amount"`       //必填 	32 	该支付工具类型所使用的金额 	10
	RealAmount  float64 `json:"real_amount"`  //渠道实际付款金额
	FundType    string  `json:"fund_type"`    //渠道所使用的资金类型,目前只在资金渠道(fund_channel)是银行卡渠道(BANKCARD)的情况下才返回该信息(DEBIT_CARD:借记卡,
	// CREDIT_CARD:信用卡,MIXED_CARD:借贷合一卡) 	DEBIT_CARD
}
type VoucherDetail struct {
	Id                         string  `json:"id"`                           //必填 	32 	券id
	Name                       string  `json:"name"`                         //必填 	64 	券名称
	Type                       string  `json:"type"`                         //必填 	32 	券类型，如： ALIPAY_FIX_VOUCHER - 全场代金券 ALIPAY_DISCOUNT_VOUCHER - 折扣券 ALIPAY_ITEM_VOUCHER - 单品优惠券 ALIPAY_CASH_VOUCHER - 现金抵价券 ALIPAY_BIZ_VOUCHER - 商家全场券 注：不排除将来新增其他类型的可能，商家接入时注意兼容性避免硬编码 	ALIPAY_FIX_VOUCHER
	Amount                     float64 `json:"amount"`                       //必填 	8 	优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute         float64 `json:"merchant_contribute"`          //可选 	8 	商家出资（特指发起交易的商家出资金额）
	OtherContribute            float64 `json:"other_contribute"`             //可选 	8 	其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo                       string  `json:"memo"`                         //可选 	256 	优惠券备注信息
	TemplateId                 string  `json:"template_id"`                  //可选 	64 	券模板id
	PurchaseBuyerContribute    string  `json:"purchase_buyer_contribute"`    //可选 	8 	如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时用户实际付款的金额
	PurchaseMerchantContribute string  `json:"purchase_merchant_contribute"` //可选 	8 	如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时商户优惠的金额
	PurchaseAntContribute      string  `json:"purchase_ant_contribute"`      //可选 	8 	如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时平台优惠的金额
}

var FundChannel = map[string]string{
	"COUPON":        "支付宝红包",
	"ALIPAYACCOUNT": "支付宝账户",
	"POINT":         "集分宝",
	"DISCOUNT":      "折扣券",
	"PCARD":         "预付卡",
	"MCARD":         "商家储值卡",
	"MDISCOUNT":     "商户优惠券",
	"MCOUPON":       "商户红包",
	"BANKCARD":      "银行卡",
} //支付渠道说明

type ApiTradeWapPay struct {
	ProductCode        string         `json:"product_code"`        //必选 	64 	销售产品码，商家和支付宝签约的产品码。 FAST_INSTANT_TRADE_PAY。
	OutTradeNo         string         `json:"out_trade_no"`        //必选	64 	商户订单号
	Subject            string         `json:"subject"`             //必选 	256 	订单标题
	TotalAmount        float64        `json:"total_amount"`        //必选 	9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 	9.00
	QuitURL            string         `json:"quit_url"`            //必选 	400 	用户付款中途退出返回商户网站的地址 	http://www.taobao.com/product/113714.html
	TimeExpire         string         `json:"time_expire"`         //可选	32 	绝对超时时间，格式为yyyy-MM-dd HH:mm:ss 	2016-12-31 10:05:01
	AuthToken          string         `json:"auth_token"`          //可选 	40 	针对用户授权接口，获取用户相关数据时，用于标识用户授权关系 	appopenBb64d181d0146481ab6a762c00714cC27
	GoodsType          string         `json:"goods_type"`          //可选 	2 	商品主类型 :0-虚拟类商品，1-实物类商品 	0
	PassbackParams     string         `json:"passback_params"`     //可选 	512 	公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。 	merchantBizType%3d3C%26merchantBizNo%3d2016010101111
	PromoParams        string         `json:"promo_params"`        //可选 	512 	优惠参数 （json 格式，仅与支付宝协商后可用） 	{"storeIdType":"1"}
	SellerID           string         `json:"seller_id"`           //可选 	28 	卖家支付宝用户 ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户 ID。不允许收款账号与付款方账号相同
	DiscountableAmount float64        `json:"discountable_amount"` //可选 	11 	参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	OperatorID         string         `json:"operator_id"`         //可选 	28 	商户操作员编号
	TerminalID         string         `json:"terminal_id"`         //可选 	32 	商户机具终端编号
	MerchantOrderNo    string         `json:"merchant_order_no"`   //可选 	32 	商户原始订单号，最大长度限制32位 	20161008001
	StoreID            string         `json:"store_id"`            //可选 	32 	商户门店编号 	NJ_001
	BusinessParams     string         `json:"business_params"`     //可选 	512 	商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，json格式 	{"data":"123"}
	ExtendParams       *ExtendParams  `json:"extend_params"`       //可选 	业务扩展参数
	QueryOptions       *[]string      `json:"query_options"`       //可选 	1024 	查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。 	["fund_bill_list","voucher_detail_list","discount_goods_detail"]
	GoodsDetail        *[]GoodsDetail `json:"goods_detail"`        //可选 	商品列表信息
}

func ApiTradeWapPayNew(OutTradeNo, Subject, QuitURL string, TotalAmount float64) ApiTradeWapPay {
	return ApiTradeWapPay{
		ProductCode: "QUICK_WAP_WAY",
		OutTradeNo:  OutTradeNo,
		Subject:     Subject,
		TotalAmount: TotalAmount,
		QuitURL:     QuitURL,
		TimeExpire:  util.TimeUtil.DateToyMdHmsSepTo(time.Now().Add(30 * time.Minute)), // 默认 30分钟付款时间 格式为yyyy-MM-dd HH:mm:ss 	2016-12-31 10:05:01
		GoodsType:   "1",
	}
} // 简单的创建

func (ApiTradeWapPay) GetMethod() string {
	return "alipay.trade.wap.pay"
}

type ApiTradePagePay struct {
	ProductCode string  `json:"product_code"` //必选 	64 	销售产品码，商家和支付宝签约的产品码。 QUICK_WAP_WAY。
	OutTradeNo  string  `json:"out_trade_no"` //必选	64 	商户订单号
	Subject     string  `json:"subject"`      //必选 	256 	订单标题
	TotalAmount float64 `json:"total_amount"` //必选 	9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 	9.00
	QuitURL     string  `json:"quit_url"`     //必选 	400 	用户付款中途退出返回商户网站的地址 	http://www.taobao.com/product/113714.html

	TimeExpire         string         `json:"time_expire"`         //可选	32 	绝对超时时间，格式为yyyy-MM-dd HH:mm:ss 	2016-12-31 10:05:01
	AuthToken          string         `json:"auth_token"`          //可选 	40 	针对用户授权接口，获取用户相关数据时，用于标识用户授权关系 	appopenBb64d181d0146481ab6a762c00714cC27
	GoodsType          string         `json:"goods_type"`          //可选 	2 	商品主类型 :0-虚拟类商品，1-实物类商品 	0
	PassbackParams     string         `json:"passback_params"`     //可选 	512 	公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。 	merchantBizType%3d3C%26merchantBizNo%3d2016010101111
	PromoParams        string         `json:"promo_params"`        //可选 	512 	优惠参数 （json 格式，仅与支付宝协商后可用） 	{"storeIdType":"1"}
	SellerID           string         `json:"seller_id"`           //可选 	28 	卖家支付宝用户 ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户 ID。不允许收款账号与付款方账号相同
	DiscountableAmount float64        `json:"discountable_amount"` //可选 	11 	参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	OperatorID         string         `json:"operator_id"`         //可选 	28 	商户操作员编号
	TerminalID         string         `json:"terminal_id"`         //可选 	32 	商户机具终端编号
	MerchantOrderNo    string         `json:"merchant_order_no"`   //可选 	32 	商户原始订单号，最大长度限制32位 	20161008001
	StoreID            string         `json:"store_id"`            //可选 	32 	商户门店编号 	NJ_001
	BusinessParams     string         `json:"business_params"`     //可选 	512 	商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，json格式 	{"data":"123"}
	ExtendParams       *ExtendParams  `json:"extend_params"`       //可选 	业务扩展参数
	QueryOptions       *[]string      `json:"query_options"`       //可选 	1024 	查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。 	["fund_bill_list","voucher_detail_list","discount_goods_detail"]
	GoodsDetail        *[]GoodsDetail `json:"goods_detail"`        //可选 	商品列表信息
	Body               string         `json:"body"`
	TimeoutExpress     string         `json:"timeout_express"`
	RoyaltyInfo        *struct {
		RoyaltyType        string `json:"royalty_type"`
		RoyaltyDetailInfos []struct {
			SerialNo         int     `json:"serial_no"`
			TransInType      string  `json:"trans_in_type"`
			BatchNo          string  `json:"batch_no"`
			OutRelationID    string  `json:"out_relation_id"`
			TransOutType     string  `json:"trans_out_type"`
			TransOut         string  `json:"trans_out"`
			TransIn          string  `json:"trans_in"`
			Amount           float64 `json:"amount"`
			Desc             string  `json:"desc"`
			AmountPercentage string  `json:"amount_percentage"`
		} `json:"royalty_detail_infos"`
	} `json:"royalty_info"`
	SubMerchant *struct {
		MerchantID   string `json:"merchant_id"`
		MerchantType string `json:"merchant_type"`
	} `json:"sub_merchant"`
	EnablePayChannels  string `json:"enable_pay_channels"`
	DisablePayChannels string `json:"disable_pay_channels"`
	QrPayMode          string `json:"qr_pay_mode"`
	QrcodeWidth        string `json:"qrcode_width"`
	SettleInfo         *struct {
		SettleDetailInfos []struct {
			TransInType      string `json:"trans_in_type"`
			SummaryDimension string `json:"summary_dimension"`
			SettleEntityID   string `json:"settle_entity_id"`
			SettleEntityType string `json:"settle_entity_type"`
			Amount           string `json:"amount"`
		} `json:"settle_detail_infos"`
	} `json:"settle_info"`
	SettlePeriodTime string `json:"settle_period_time"`
	InvoiceInfo      *struct {
		KeyInfo struct {
			IsSupportInvoice    bool   `json:"is_support_invoice"`
			InvoiceMerchantName string `json:"invoice_merchant_name"`
			TaxNum              string `json:"tax_num"`
		} `json:"key_info"`
		Details string `json:"details"`
	} `json:"invoice_info"`
	AgreementSignParams *AgreementSignParams `json:"agreement_sign_params"` //签约参数，支付后签约场景使用
	IntegrationType     string               `json:"integration_type"`
	RequestFromURL      string               `json:"request_from_url"`
	ExtUserInfo         *ExtUserInfo         `json:"ext_user_info"`
}

func (ApiTradePagePay) GetMethod() string {
	return "alipay.trade.page.pay"
}
func ApiTradePagePayNew(OutTradeNo, Subject, QuitURL string, TotalAmount float64) ApiTradePagePay {
	return ApiTradePagePay{
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		OutTradeNo:  OutTradeNo,
		Subject:     Subject,
		TotalAmount: TotalAmount,
		QuitURL:     QuitURL,
		TimeExpire:  util.TimeUtil.DateToyMdHmsSepTo(time.Now().Add(30 * time.Minute)), // 默认 30分钟付款时间 格式为yyyy-MM-dd HH:mm:ss 	2016-12-31 10:05:01
		GoodsType:   "1",
	}
} // 简单的创建

type ApiTradeRefund struct {
	OutTradeNo              string                        `json:"out_trade_no"`              // 	String 	必选 	64 	商户订单号。64 个字符以内的大小，仅支持字母、数字、下划线。需保证该参数在商户端不重复。 	20150320010101001
	GoodsDetail             *[]GoodsDetail                `json:"goods_detail"`              //		可选 	商品列表信息
	TradeNo                 string                        `json:"trade_no"`                  // 	String 	特殊可选 	64 	支付宝交易号，和商户订单号 out_trade_no 不能同时为空。 	2014112611001004680073956707
	RefundAmount            float64                       `json:"refund_amount"`             // 	Price 	必选 	11 	需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数 	200.12
	RefundCurrency          string                        `json:"refund_currency"`           // 	String 	可选 	8 	订单退款币种信息。支持英镑：GBP、港币：HKD、美元：USD、新加坡元：SGD、日元：JPY、加拿大元：CAD、澳元：AUD、欧元：EUR、新西兰元：NZD、韩元：KRW、泰铢：THB、瑞士法郎：CHF、瑞典克朗：SEK、丹麦克朗：DKK、挪威克朗：NOK、马来西亚林吉特：MYR、印尼卢比：IDR、菲律宾比索：PHP、毛里求斯卢比：MUR、以色列新谢克尔：ILS、斯里兰卡卢比：LKR、俄罗斯卢布：RUB、阿联酋迪拉姆：AED、捷克克朗：CZK、南非兰特：ZAR、人民币：CNY 	USD
	RefundReason            string                        `json:"refund_reason"`             // 	String 	可选 	256 	退款原因说明，商家自定义。 	正常退款
	OutRequestNo            string                        `json:"out_request_no"`            // 	String 	可选 	64 	标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。 	HZ01RF001
	OperatorID              string                        `json:"operator_id"`               // 	String 	可选 	30 	商户的操作员编号 	OP001
	StoreID                 string                        `json:"store_id"`                  // 	String 	可选 	32 	商户门店编号，由商家自定义。需保证当前商户下唯一。 	NJ_S_001
	TerminalID              string                        `json:"terminal_id"`               // 	String 	可选 	32 	商户的终端编号 	NJ_T_001
	RefundRoyaltyParameters *OpenApiRoyaltyDetailInfoPojo `json:"refund_royalty_parameters"` // 	OpenApiRoyaltyDetailInfoPojo[] 	可选 		退分账明细信息。注：1.当面付无需传入退分账明细，系统自动按退款金额与订单金额的比率，从收款方和分账收入方退款，不支持指定退款金额与退款方。	2.电脑网站支付，手机 APP 支付，手机网站支付产品，须在退款请求中明确是否退分账，从哪个分账收入方退，退多少分账金额；如不明确，默认从收款方退款，收款方余额不足退款失败。不支持系统按比率退款。
	OrgPid                  string                        `json:"org_pid"`                   // 	String 	可选 	16 	银行间联模式下有用，其它场景请不要使用；双联通过该参数指定需要退款的交易所属收单机构的pid; 	2088101117952222
	QueryOptions            *[]string                     `json:"query_options"`             // 	String[] 	可选 	1024 	查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。支持：refund_detail_item_list：退款使用的资金渠道。 	refund_detail_item_list
} //alipay.trade.refund(统一收单交易退款接口)
/*
当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，
卖家可以通过退款接口将支付款退还给买家，
支付宝将在收到退款请求并且验证成功之后，
按照退款规则将支付款按原路退到买家帐号上。
交易超过约定时间（签约时设置的可退款时间）的订单无法进行退款
支付宝退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。
一笔退款失败后重新提交，要采用原来的退款单号。
总退款金额不能超过用户实际支付金额
*/
func (ApiTradeRefund) GetMethod() string {
	return "alipay.trade.refund"
}
func ApiTradeRefundNew(OutTradeNo, TradeNo, RefundReason string, RefundAmount float64) ApiTradeRefund {
	return ApiTradeRefund{
		RefundReason: RefundReason,
		TradeNo:      TradeNo,
		OutTradeNo:   OutTradeNo,
		RefundAmount: RefundAmount,
	}
} // 简单的创建
type OpenApiRoyaltyDetailInfoPojo struct {
	RoyaltyType      string  `json:"royalty_type"`      // 	String 	可选 	32 	分账类型.普通分账为：transfer;补差为：replenish;为空默认为分账transfer; 	transfer
	TransOut         string  `json:"trans_out"`         // 	String 	可选 	16 	支出方账户。如果支出方账户类型为userId，本参数为支出方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果支出方类型为loginName，本参数为支出方的支付宝登录号； 	2088101126765726
	TransOutType     string  `json:"trans_out_type"`    // 	String 	可选 	64 	支出方账户类型。userId表示是支付宝账号对应的支付宝唯一用户号;loginName表示是支付宝登录号； 	userId
	TransInType      string  `json:"trans_in_type"`     // 	String 	可选 	64 	收入方账户类型。userId表示是支付宝账号对应的支付宝唯一用户号;cardAliasNo表示是卡编号;loginName表示是支付宝登录号； 	userId
	TransIn          string  `json:"trans_in"`          // 	String 	必填 	16 	收入方账户。如果收入方账户类型为userId，本参数为收入方的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；如果收入方类型为cardAliasNo，本参数为收入方在支付宝绑定的卡编号；如果收入方类型为loginName，本参数为收入方的支付宝登录号； 	2088101126708402
	Amount           float64 `json:"amount"`            // 	Price 	可选 	9 	分账的金额，单位为元 	0.1
	AmountPercentage int     `json:"amount_percentage"` // 	Number 	可选 	3 	分账信息中分账百分比。取值范围为大于0，少于或等于100的整数。 	100
	Desc             string  `json:"desc"`              // 	String 	可选 	1000 	分账描述 	分账给2088101126708402
} //退分账明细信息

type ApiTradePageRefund struct {
	ProductCode string  `json:"product_code"` //必选 	64 	销售产品码，商家和支付宝签约的产品码。 QUICK_WAP_WAY。
	OutTradeNo  string  `json:"out_trade_no"` //必选	64 	商户订单号
	Subject     string  `json:"subject"`      //必选 	256 	订单标题
	TotalAmount float64 `json:"total_amount"` //必选 	9 	订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 	9.00
	QuitURL     string  `json:"quit_url"`     //必选 	400 	用户付款中途退出返回商户网站的地址 	http://www.taobao.com/product/113714.html

	TimeExpire         string         `json:"time_expire"`         //可选	32 	绝对超时时间，格式为yyyy-MM-dd HH:mm:ss 	2016-12-31 10:05:01
	AuthToken          string         `json:"auth_token"`          //可选 	40 	针对用户授权接口，获取用户相关数据时，用于标识用户授权关系 	appopenBb64d181d0146481ab6a762c00714cC27
	GoodsType          string         `json:"goods_type"`          //可选 	2 	商品主类型 :0-虚拟类商品，1-实物类商品 	0
	PassbackParams     string         `json:"passback_params"`     //可选 	512 	公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在同步返回（包括跳转回商户网站）和异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝。 	merchantBizType%3d3C%26merchantBizNo%3d2016010101111
	PromoParams        string         `json:"promo_params"`        //可选 	512 	优惠参数 （json 格式，仅与支付宝协商后可用） 	{"storeIdType":"1"}
	SellerID           string         `json:"seller_id"`           //可选 	28 	卖家支付宝用户 ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户 ID。不允许收款账号与付款方账号相同
	DiscountableAmount float64        `json:"discountable_amount"` //可选 	11 	参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	OperatorID         string         `json:"operator_id"`         //可选 	28 	商户操作员编号
	TerminalID         string         `json:"terminal_id"`         //可选 	32 	商户机具终端编号
	MerchantOrderNo    string         `json:"merchant_order_no"`   //可选 	32 	商户原始订单号，最大长度限制32位 	20161008001
	StoreID            string         `json:"store_id"`            //可选 	32 	商户门店编号 	NJ_001
	BusinessParams     string         `json:"business_params"`     //可选 	512 	商户传入业务信息，具体值要和支付宝约定，应用于安全，营销等参数直传场景，json格式 	{"data":"123"}
	ExtendParams       *ExtendParams  `json:"extend_params"`       //可选 	业务扩展参数
	QueryOptions       *[]string      `json:"query_options"`       //可选 	1024 	查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。 	["fund_bill_list","voucher_detail_list","discount_goods_detail"]
	GoodsDetail        *[]GoodsDetail `json:"goods_detail"`        //可选 	商品列表信息
	Body               string         `json:"body"`
	TimeoutExpress     string         `json:"timeout_express"`
	RoyaltyInfo        *struct {
		RoyaltyType        string `json:"royalty_type"`
		RoyaltyDetailInfos []struct {
			SerialNo         int     `json:"serial_no"`
			TransInType      string  `json:"trans_in_type"`
			BatchNo          string  `json:"batch_no"`
			OutRelationID    string  `json:"out_relation_id"`
			TransOutType     string  `json:"trans_out_type"`
			TransOut         string  `json:"trans_out"`
			TransIn          string  `json:"trans_in"`
			Amount           float64 `json:"amount"`
			Desc             string  `json:"desc"`
			AmountPercentage string  `json:"amount_percentage"`
		} `json:"royalty_detail_infos"`
	} `json:"royalty_info"`
	SubMerchant *struct {
		MerchantID   string `json:"merchant_id"`
		MerchantType string `json:"merchant_type"`
	} `json:"sub_merchant"`
	EnablePayChannels  string `json:"enable_pay_channels"`
	DisablePayChannels string `json:"disable_pay_channels"`
	QrPayMode          string `json:"qr_pay_mode"`
	QrcodeWidth        string `json:"qrcode_width"`
	SettleInfo         *struct {
		SettleDetailInfos []struct {
			TransInType      string `json:"trans_in_type"`
			SummaryDimension string `json:"summary_dimension"`
			SettleEntityID   string `json:"settle_entity_id"`
			SettleEntityType string `json:"settle_entity_type"`
			Amount           string `json:"amount"`
		} `json:"settle_detail_infos"`
	} `json:"settle_info"`
	SettlePeriodTime string `json:"settle_period_time"`
	InvoiceInfo      *struct {
		KeyInfo struct {
			IsSupportInvoice    bool   `json:"is_support_invoice"`
			InvoiceMerchantName string `json:"invoice_merchant_name"`
			TaxNum              string `json:"tax_num"`
		} `json:"key_info"`
		Details string `json:"details"`
	} `json:"invoice_info"`
	AgreementSignParams *AgreementSignParams `json:"agreement_sign_params"` //签约参数，支付后签约场景使用
	IntegrationType     string               `json:"integration_type"`
	RequestFromURL      string               `json:"request_from_url"`
	ExtUserInfo         *ExtUserInfo         `json:"ext_user_info"`
} //统一收单退款页面接口
/*
当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，
卖家可以通过退款页面接口将支付款退还给买家，
支付宝将在收到退款请求并且验证成功之后，
按照退款规则将支付款按原路退到买家帐号上。
目前该接口用于信用退款场景，通过biz_type指定信用退款。
支付宝页面会提示用户退款成功或失败，
退款处理完成后支付宝回跳到商户请求指定的回跳地址页面。
*/
func (ApiTradePageRefund) GetMethod() string {
	return "alipay.trade.page.refund"
}
func ApiTradePageRefundNew(OutTradeNo, Subject, QuitURL string, TotalAmount float64) ApiTradePageRefund {
	return ApiTradePageRefund{
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		OutTradeNo:  OutTradeNo,
		Subject:     Subject,
		TotalAmount: TotalAmount,
		QuitURL:     QuitURL,
		TimeExpire:  util.TimeUtil.DateToyMdHmsSepTo(time.Now().Add(30 * time.Minute)), // 默认 30分钟付款时间 格式为yyyy-MM-dd HH:mm:ss 	2016-12-31 10:05:01
		GoodsType:   "1",
	}
} // 简单的创建

type ApiTradeQuery struct {
	OutTradeNo   string   `json:"out_trade_no"`  //String 	特殊可选 	64 	订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no 	20150320010101001
	TradeNo      string   `json:"trade_no"`      // 	String 	特殊可选 	64 	支付宝交易号，和商户订单号不能同时为空 	2014112611001004680 073956707
	OrgPid       string   `json:"org_pid"`       // 	String 	可选 	16 	银行间联模式下有用，其它场景请不要使用；双联通过该参数指定需要查询的交易所属收单机构的pid; 	2088101117952222
	QueryOptions []string `json:"query_options"` // 	String[] 	可选 	1024 	查询选项，商户传入该参数可定制本接口同步响应额外返回的信息字段，数组格式。支持枚举如下：trade_settle_info：返回的交易结算信息，包含分账、补差等信息。fund_bill_list：交易支付使用的资金渠道。 	trade_settle_info

} //统一收单线下交易查询
/*
该接口提供所有支付宝支付订单的查询，商户可以通过该接口主动查询订单状态，
完成下一步的业务逻辑。 需要调用查询接口的情况：
当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；
调用支付接口后，返回系统错误或未知交易状态情况；
调用alipay.trade.pay，返回INPROCESS的状态；
调用alipay.trade.cancel之前，需确认支付状态；
*/
func (ApiTradeQuery) GetMethod() string {
	return "alipay.trade.query"
}
func ApiTradeQueryNew(OutTradeNo, TradeNo string) ApiTradeQuery {
	return ApiTradeQuery{
		OutTradeNo: OutTradeNo,
		TradeNo:    TradeNo,
	}
} // 简单的创建
type ExtUserInfo struct {
	Name          string `json:"name"`            //姓名 注： need_check_info=T时该参数才有效
	Mobile        string `json:"mobile"`          //手机号
	CertType      string `json:"cert_type"`       //IDENTITY_CARD 身份证：IDENTITY_CARD、护照：PASSPORT、军官证：OFFICER_CARD、士兵证：SOLDIER_CARD、户口本：HOKOU等。如有其它类型需要支持，请与蚂蚁金服工作人员联系。
	CertNo        string `json:"cert_no"`         //证件号 注：need_check_info=T时该参数才有效
	MinAge        string `json:"min_age"`         //18 允许的最小买家年龄，买家年龄必须大于等于所传数值 注：1. need_check_info=T时该参数才有效2. min_age为整数，必须大于等于0
	FixBuyer      string `json:"fix_buyer"`       //F 是否强制校验付款人身份信息 T:强制校验，F：不强制
	NeedCheckInfo string `json:"need_check_info"` //F 是否强制校验身份信息 T:强制校验，F：不强制 F
} //外部指定买家
type AgreementSignParams struct {
	PersonalProductCode string `json:"personal_product_code"` // 	String 	必填 	64 	个人签约产品码，商户和支付宝签约时确定。 	GENERAL_WITHHOLDING_P
	SignScene           string `json:"sign_scene"`            // 	String 	可选 	64 	协议签约场景，商户和支付宝签约时确定。当传入商户签约号external_agreement_no时，场景不能为默认值DEFAULT|DEFAULT。 	INDUSTRY|CARRENTAL
	ExternalAgreementNo string `json:"external_agreement_no"` // 	String 	可选 	32 	商户签约号，代扣协议中标示用户的唯一签约号（确保在商户系统中唯一）。格式规则：支持大写小写字母和数字，最长32位。商户系统按需传入，如果同一用户在同一产品码、同一签约场景下，签订了多份代扣协议，那么需要指定并传入该值。 	test
	ExternalLogonID     string `json:"external_logon_id"`     // 	String 	可选 	100 	用户在商户网站的登录账号，用于在签约页面展示，如果为空，则不展示 	13852852877
	SignValidityPeriod  string `json:"sign_validity_period"`  // 	String 	可选 	8 	当前用户签约请求的协议有效周期。整形数字加上时间单位的协议有效期，从发起签约请求的时间开始算起。目前支持的时间单位：1. d：天	2. m：月	如果未传入，默认为长期有效。 	2m
	ThirdPartyType      string `json:"third_party_type"`      // 	String 	可选 	32 	签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。取值范围：1. PARTNER（平台商户）;2. MERCHANT（集团商户），集团下子商户可共享用户签约内容;默认为PARTNER。 	PARTNER
	BuckleAppID         string `json:"buckle_app_id"`         // 	String 	可选 	64 	商户在芝麻端申请的appId 	1001164
	BuckleMerchantID    string `json:"buckle_merchant_id"`    // 	String 	可选 	64 	商户在芝麻端申请的merchantId 	268820000000414397785
	PromoParams         string `json:"promo_params"`          // 	String 	可选 	512 	签约营销参数，此值为json格式；具体的key需与营销约定 	{"key","value"}
} //签约参数，支付后签约场景使用
type ExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id"` //64 	系统商编号 该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID 	2088511833207846
	HbFqNum              string `json:"hb_fq_num"`               //5 	使用花呗分期要进行的分期数 	3
	HbFqSellerPercent    string `json:"hb_fq_seller_percent"`    //3 	使用花呗分期需要卖家承担的手续费比例的百分值，传入100代表100% 	100
	IndustryRefluxInfo   string `json:"industry_reflux_info"`    //512 	行业数据回流信息, 详见：地铁支付接口参数补充说明 	{\"scene_code\":\"metro_tradeorder\", \"channel\":\"xxxx\",\"scene_data\":{\"asset_name\":\"ALIPAY\"}}
	CardType             string `json:"card_type"`               //32 	卡类型 	S0JP0000
} //业务扩展参数
type GoodsDetail struct {
	GoodsID   string  `json:"goods_id"`   //必填 	32 	商品的编号 	apple-01
	GoodsName string  `json:"goods_name"` //必填 	256 	商品名称
	Quantity  int     `json:"quantity"`   //必填 	10 	商品数量
	Price     float64 `json:"price"`      //必填 	9 	商品单价，单位为元 	2000

	AlipayGoodsID  string `json:"alipay_goods_id"` //可选 	32 	支付宝定义的统一商品编号 	20010001
	GoodsCategory  string `json:"goods_category"`  //可选 	24 	商品类目 	34543238
	CategoriesTree string `json:"categories_tree"` //可选 	128 	商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割 	124868003|126232002|126252004
	Body           string `json:"body"`            //可选 	1000 	商品描述信息 	特价手机
	ShowURL        string `json:"show_url"`        //可选 	400 	商品的展示地址 	http://www.alipay.com/xxx.jpg
} //商品列表信息
