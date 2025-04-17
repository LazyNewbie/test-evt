<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Exads\Common\Schema\Events;

/**
 *
 * Event{
 * Date:                      eventData["aa"].(int32),             date_int
 * Language:                  eventData["ad"].(int32),             idlanguage
 * Advertiser:                eventData["af"].(int64),             idadvertiser
 * Publisher:                 eventData["ag"].(int64),             idpublisher
 * Site:                      eventData["ah"].(int64),             idsite
 * Zone:                      eventData["ai"].(int64),             idzone
 * Category:                  eventData["aj"].(int32),             idcategory
 * Campaign:                  eventData["ak"].(int64),             idcampaign
 * Variation:                 eventData["al"].(int64),             idvariation
 * OS:                        eventData["am"].(int32),             idos
 * Browser:                   eventData["an"].(int32),             idbrowser
 * Device:                    eventData["ao"].(int32),             iddevice
 * Carrier:                   eventData["ap"].(int32),             idcarrier
 * ISP:                       eventData["aq"].(int32),             idisp
 * Sub:                       eventData["ar"].(int64),             sub
 * Sub2:                      eventData["dc"].(int64),             sub2
 * Sub3:                      eventData["dd"].(int64),             sub3
 * PublisherRate:             eventData["at"].(int32),             publisher_rate
 * AdType:                    eventData["ay"].(int32),             ad_type
 * Network:                   eventData["ba"].(int32),             idnetwork
 * ImpressionsStatus:         eventData["bb"].(int32),             impression_status
 * ImpressionInvalidReason:   eventData["bc"].(int32),             impression_invalid_reason
 * AdBlock:                   eventData["bd"].(int32),             adblock
 * Offer:                     eventData["be"].(int64),             idoffer
 * LandingPage:               eventData["bf"].(int64),             idlanding_page
 * Position:                  eventData["bg"].(int32),             position
 * RegionGeonameID:           eventData["bh"].(int64),             region_geoname_id
 * Partner:                   eventData["bi"].(int32),             idpartner
 * ClickStatus:               eventData["bj"].(int32),             click_status
 * RtbToalTime:               eventData["bn"].(int64),             rtb_total_time
 * RtbRequestType:            eventData["bt"].(int32),             rtb_request_type
 * Type:                      eventData["bu"].(int32),             event_type
 * CityGeonameID:             eventData["cp"].(int32),             city_geoname_id
 * ClickAllInvalidReasons:    eventData["bl"].(int32),             click_all_invalid_reasons
 * ClickX:                    eventData["ct"].(int32),             clickx
 * ClickY:                    eventData["cu"].(int32),             clicky
 * ErrorCode:                 eventData["ch"].(int32),             errorcode
 * HTTPCode:                  eventData["ck"].(int32),             http_code
 * HTTPS:                     eventData["ca"].(int32),             https
 * CarrierSource:             eventData["cl"].(int32),             idcarrier_source
 * Datacenter:                eventData["cr"].(int32),             iddatacenter
 * ProductCategory:           eventData["by"].(int32),             idproduct_category
 * IncognitoMode:             eventData["cn"].(int32),             incognito_mode
 * TrustedProxy:              eventData["cq"].(int32),             is_trusted_proxy
 * NoScript:                  eventData["cm"].(int32),             noscript
 * PricingModel:              eventData["bz"].(int32),             pricing_model
 * TestMode:                  eventData["co"].(int32),             test_mode
 * VideoEvent:                eventData["cg"].(int32),             video_event
 * WpnEvent:                  eventData["da"].(int32),             wpn_event
 * EmailStatus:               eventData["db"].(int32),             email_status
 * IdClientHints:             eventData["eu"].(int32),             client_hints_usage
 * Time:                      eventData["ab"].(int64),             time
 * Value:                     eventData["as"].(float64),           value
 * Exchange2Publisher:        eventData["aw"].(float64),           exchange_to_publisher
 * Exchange2Network:          eventData["ax"].(float64),           exchange_to_network
 * ExchangeEUR2USD:           eventData["cz"].(float64),           eur_to_usd
 * RtbContentSize:            eventData["bo"].(float64),           rtb_content_size
 * RtbValue:                  eventData["bs"].(float64),           rtb_value
 * ExecutionTime:             eventData["cf"].(float64),           execution_time
 * ConversionValueEUR:        eventData["dr"].(float64),           conversion_value_eur
 * ConversionValueUSD:        eventData["ds"].(float64),           conversion_value_usd
 * Country:                   eventData["ac"].(string),            country
 * SiteHostname:              eventData["ae"].(string),            site_hostname
 * PublisherCurrency:         eventData["au"].(string),            publisher_currency
 * CampaignCurrency:          eventData["av"].(string),            campaign_currency
 * AdFormat:                  eventData["az"].(string),            format
 * ClickInvalidReason:        eventData["bk"].(string),            click_invalid_reason
 * Goal:                      eventData["bm"].(string),            idgoal
 * BidStatus:                 eventData["bp"].(string),            bid_status
 * ErrorReason:               eventData["bq"].(string),            error_reason
 * AdblockVersion:            eventData["cs"].(string),            adblock_ver
 * BidID:                     eventData["cj"].(string),            bid_id
 * ContainerResolution:       eventData["cv"].(string),            container_resolution
 * CountryOriginal:           eventData["bw"].(string),            country_original
 * HTTPXForwarderFor:         eventData["bx"].(string),            http_x_forwarded_for
 * ImpID:                     eventData["ci"].(string),            imp_id
 * NetacuityProxyDescription: eventData["cx"].(string),            netacuity_proxy_description
 * NetacuityProxyType:        eventData["cy"].(string),            netacuity_proxy_type
 * RefererSiteHostname:       eventData["ce"].(string),            referer_site_hostname
 * ScreenResolution:          eventData["cb"].(string),            screen_resolution
 * URLHash:                   eventData["cd"].(string),            url_hash
 * UserID:                    eventData["cc"].(string),            user_id
 * UserIP:                    eventData["bv"].(string),            user_ip
 * ScriptType:                eventData["de"].(string),            script_type
 * ScriptName:                eventData["df"].(string),            script_name
 * ScriptVer:                 eventData["dg"].(string),            script_ver
 * RtbInventoryId:            eventData["dx"].(string),            rtb_inventory_id
 * RtbInventoryName:          eventData["dy"].(string),            rtb_inventory_name
 * AdServingHost:             eventData["es"].(string),            ad_serving_host
 * CustomValue1:              eventData["ev"].(string),            custom_value_1
 * CustomValue2:              eventData["ew"].(string),            custom_value_2
 * CustomValue3:              eventData["ex"].(string),            custom_value_3
 * BrowserVer:                eventData["dh"].(int32),             browser_ver
 * EmailISP:                  eventData["do"].(int32),             email_isp
 * VR:                        eventData["di"].(int32),             vr
 * Internal:                  eventData["dj"].(int32),             internal
 * Sampling:                  eventData["dk"].(int32),             sampling
 * NoFunds:                   eventData["dl"].(int32),             no_funds
 * RtbQueryStatus:            eventData["br"].(bool),              rtb_query_status
 * ResponsiveLayout:          eventData["dm"].(int32),             idresponsive_layout
 * MultiZone:                 eventData["dn"].(int32),             idmultizone
 * ScriptDetail:              eventData["dp"].(int32),             script_detail
 * ScriptVersion:             eventData["dq"].(int32),             script_version
 * IdadvertiserAdType:        eventData["dt"].(int32),             idadvertiser_ad_type
 * Proxy:                     eventData["du"].(int32),             idproxy
 * AdRequestResult:           eventData["dv"].(int32),             ad_request_result
 * TrafficType:               eventData["dw"].(int32),             idnetwork_type
 * TrafficSource:             eventData["eq"].(int32),             idtraffic_source
 * FallbackPosition:          eventData["er"].(int32),             fallback_position
 * IdCampaignType:            eventData["dz"].(int32),             idcampaign_type
 * RefreshIteration:          eventData["ea"].(int32),             refresh_iteration
 * KeyValue1:                 eventData["eb"].(int32),             key_value_1
 * KeyValue2:                 eventData["ec"].(int32),             key_value_2
 * KeyValue3:                 eventData["ed"].(int32),             key_value_3
 * KeyValue4:                 eventData["ee"].(int32),             key_value_4
 * KeyValue5:                 eventData["ef"].(int32),             key_value_5
 * KeyValue6:                 eventData["eg"].(int32),             key_value_6
 * KeyValue7:                 eventData["eh"].(int32),             key_value_7
 * KeyValue8:                 eventData["ei"].(int32),             key_value_8
 * KeyValue9:                 eventData["ej"].(int32),             key_value_9
 * KeyValue10:                eventData["ek"].(int32),             key_value_10
 * IdRtbErrorReason:          eventData["et"].(int32),             idrtb_error_reason
 * CdnUsage:                  eventData["el"].(int64),             cdn_usage
 * BidShadingValue:           eventData["em"].(float32),           bid_shading_value
 * SecondBestValue:           eventData["en"].(float32),           second_best_value
 * AuctionFloor:              eventData["eo"].(float32),           auction_floor
 * AuctionCeil:               eventData["ep"].(float32),           auction_ceil
 * Raw:                       eventData,
 * }
 *
 *
 *
 *
 *
 * Array
 * (
 * [impression_all_invalid_reasons] => "cw"
 * )
 *
 *
 *
 * The greeting service definition.
 */
class AdEventsServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Sends a greeting
     * @param \Exads\Common\Schema\Events\AdEvent $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Enqueue(\Exads\Common\Schema\Events\AdEvent $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/exads.schema.events.AdEventsService/Enqueue',
        $argument,
        ['\Exads\Common\Schema\Events\AdEventResponse', 'decode'],
        $metadata, $options);
    }

}
