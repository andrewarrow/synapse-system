package chrome

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
)

type Cookie struct {
	Name   string
	Value  string
	Domain string
}

func Run(email string) {
	url := "https://app.slack.com/client"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res []*cdp.Node
	cookies := []Cookie{}

	c = Cookie{"activitySession_T05G01UF8DC","{}"}
c = Cookie{"ChimeRegionStorage-chimeRegionCache","{"regions":"us-west-1","timestamp":1689101468616,"version":1}"}
c = Cookie{"client-ia-history","[{"ts":1689101463275,"teamId":"T05G01UF8DC","experimentClsMoreHistoryGroupTreatment":true,"primaryView":"C05G01X790W","position":1}]"}
c = Cookie{"encryption-NbQ9EAOWHQ2-key-version","5"}
c = Cookie{"endpoints-v1-T05G01UF8DC","{"socketUrl":{"url":"wss://wss-primary.slack.com/","fallbackUrl":"wss://wss-backup.slack.com/","ttlSeconds":604800,"context":"T05G01UF8DC-5","expiry":1689706261115}}"}
c = Cookie{"localConfig_v2","{"teams":{"T05G01UF8DC":{"id":"T05G01UF8DC","name":"many.pw","url":"https://manypw.slack.com/","domain":"manypw","token":"xoxc-5544062518454-5547833135317-5556119155910-ca059a4f657c8d2739776ad013c068c144e975e11fe8fb7916af6537fce02f0b","user_locale":"en-US","user_id":"U05G3QH3Z9B","is_unified_user_client_enabled":false,"icon":{"image_68":"https://avatars.slack-edge.com/2023-07-09/5574444255424_fb14f0b756c2347a0645_68.png","image_88":"https://avatars.slack-edge.com/2023-07-09/5574444255424_fb14f0b756c2347a0645_88.png","image_default":false},"lastRoute":{"routeName":"ROUTE_ENTITY","params":{"teamId":"T05G01UF8DC","entityId":"C05G01X790W"}},"calculatedSidebarWidth":220,"channelSidebarBackground":"#19171D","teamSwitcherBackground":"rgb(5,3,9)","textColor":"#D1D2D3","topNavBackground":"#121016","topNavTextColor":"#D1D2D3","versionDataTs":1689097659}},"prevTeams":{},"orderedTeamIds":["T05G01UF8DC"],"pendingAuthTeams":{},"lc":"1689101455","lastFetched":"1689101461","lastActiveTeamId":"T05G01UF8DC","canAccessClientV2":null,"lastColdBootTs":1689101463036,"composerUpArrowShortcutBehavior":1}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::accessibilityPersistent","{"hasKeyboardFocus":false}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::customStatus","{}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::draftContactUnfurls","{}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::draftOffline","{"pendingCreates":[],"pendingDeletes":[],"pendingUpdates":[]}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::drafts","{"channelIds":["D05G9627LS0"],"channelDrafts":{"D05G9627LS0":{"id":"D05G9627LS0","ops":[{"insert":"hey\n"}],"range":{"index":0,"length":0},"date_created":1688956084,"last_updated":1688956084397.789,"broadcast":false}},"withheldDraftIds":[],"unifiedDrafts":{"D05G9627LS0":{"is_from_composer":false,"date_created":1688956084,"last_updated":1688956084397.789,"destinations":[{"channel_id":"D05G9627LS0"}],"ops":[{"insert":"hey\n"}],"file_ids":[],"cursor_index":0,"client_draft_id":"D05G9627LS0","id":"Dr05G0322XAA","last_updated_ts":1688956084.397789,"date_scheduled":0,"temp_attached_draft_unique_id":"c401b018-aca4-4c6e-ad77-28c5f020a934"}},"isEditingDraftPage":false,"selectedDraftIds":[]}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::draftUnfurls","{}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::expandables","{}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::filePermissions","{}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::persistedApiCalls","{}"}
c = Cookie{"persist-v1::T05G01UF8DC::U05G3QH3Z9B::recentlyJoinedChannels","{}"}
c = Cookie{"persist:slack-persistence-method-version-client-T05G01UF8DC","3"}
c = Cookie{"reducedSignInWorkflowExperiment","{"isOn":true,"lastActiveTeamId":"T05G01UF8DC"}"}
c = Cookie{"server-options-canned_thread_metadata_by_name",""{}""}
c = Cookie{"server-options-company_admin_min_level","10"}
c = Cookie{"server-options-company_member_lookup_large_company_ids","[]"}
c = Cookie{"server-options-copy_folder_max_threads","100"}
c = Cookie{"server-options-css_autoreload","false"}
c = Cookie{"server-options-css_autoreload_frequency","1000"}
c = Cookie{"server-options-default_slide_layout_canned_thread_names","["default_slide_layouts_titles","default_slide_layouts_text","default_slide_layouts_data","default_slide_layouts_media","default_slide_layouts_diagrams"]"}
c = Cookie{"server-options-diagnostic_report_max_mutations","200"}
c = Cookie{"server-options-edge_resources_cookie",""edge_resources""}
c = Cookie{"server-options-edge_resources_cookie_latest_val",""LATEST""}
c = Cookie{"server-options-edit_trace_enabled_threads","[]"}
c = Cookie{"server-options-edit_trace_max_sections","10"}
c = Cookie{"server-options-edit_trace_sample_percentage","1"}
c = Cookie{"server-options-elements_explorer_promoted_element_ids","[]"}
c = Cookie{"server-options-elements_explorer_template_data",""{}""}
c = Cookie{"server-options-elements_url_protocol_allowlist","["http","https","jamfselfservice","keybase","message","notes","slack","slackmdm","slackintune","slackdebug","slackbeta","slackprototype","smb","xmpp","vnc","goodtask3","sfanalytics","adir","alchemy","core","efin","fusionx","hdfs","hello","ilog","its-conductor","livability","nfa","phantom","x-phantom","rdar","repair","rxcloud","snr","snrx","swbd","thegrove","unib","mailto","tel"]"}
c = Cookie{"server-options-emoji_cdn_hash_activity",""mVFXFYpnQ60Ih7iriSz-zQ""}
c = Cookie{"server-options-emoji_cdn_hash_flags",""5clQGBH57PH88rYb7r-tTA""}
c = Cookie{"server-options-emoji_cdn_hash_food",""o4MdrAvsqGo_-WBHF3IvcA""}
c = Cookie{"server-options-emoji_cdn_hash_nature",""sSMDHjSRucwm7JS9jLbrxw""}
c = Cookie{"server-options-emoji_cdn_hash_objects",""_tcWulUWYzPWhfPlnqR9qw""}
c = Cookie{"server-options-emoji_cdn_hash_people",""EULlNowX13KEPaEo97xrAg""}
c = Cookie{"server-options-emoji_cdn_hash_symbols",""ma1lFmrVWCdhHezny0sizA""}
c = Cookie{"server-options-emoji_cdn_hash_travel",""oWI3NrRu3p-Bne5wUeR5GQ""}
c = Cookie{"server-options-empty_document_data_hash",""5ba93c9db0cff93f52b521d7420e43f6eda2784f""}
c = Cookie{"server-options-enable_inbound_sharing_restriction_check","false"}
c = Cookie{"server-options-export_company_max_members","1000"}
c = Cookie{"server-options-export_csv_filtered_company_members","false"}
c = Cookie{"server-options-external_sharing_support_article_link",""https://help.salesforce.com/s/articleView?id=sf.quip_restricted_sharing.htm""}
c = Cookie{"server-options-featured_template_data",""[]""}
c = Cookie{"server-options-folder_content_over_limit_company_list","["NQHAcAV4DT7","XEQAcA2S22G","ZWBAcAaVMgw","fBJAcAi1ucG","EYHAcAE74b0","dLQAcArjo4o","fRTAcAKiFYF","XbaAcAr82Xu","bMEAcAiFppE","XSCAcAaV0Z6","KIWAcAdGdH5","VeB9cAKFALF","CGZAcAecPQc","EaDAcAl3CKD","TXZAcAZ6Xs4","LeGAcANGSf2","HeUAcAwLAk9","IOO9cAA1JUW","YbLAcAJ8zEl","KAWAcAsdD66","FKFAcA6xsmB","OBBAcACQTcc","bAQ9cAkaQjr","aEC9cAYenVt","RHVAcAjOwAv","RAFAcAaANNt","ScIAcAfUvYI","bbYAcAcyC1T","WaQAcAuhMhm","fPGAcAqvqys","MDCAcAE66UW","WXcAcAZaqwl","RcYAcANHiZf","WeUAcAH8s5C","NPYAcA1vd4T","JLZAcAi8Siw","TdQAcAsnZXq","DGWAcAqMV0J","JYPAcABbByD","eBMAcAg0KDi","IKYAcAoFyCb","CBMAcAopDzi","TNGAcAjeKiq","FHdAcAMs8rW","YDeAcAA43dv","ROTAcAuEKAk","UESAcAo80Jv","TPeAcAr3PLb","WYFAcAY06Qz","WXfAcAVUIIg","RYOAcAT66cO","dBEAcA7G1p1","CDLAcAPyU9E","MXIAcAqBDB8","XaJAcA6fRKO","WBXAcAs6H8q","YUWAcADVXjx","NZIAcAJtbTt","dZNAcABEJVG","GHZAcAT9hR9","XAeAcAmEwCS","SZdAcAaZRHN","eZGAcANOdqy","bJbAcA9aRFY","LZXAcAeEcgH","UUfAcAhl4A5","DFEAcAn5128","BVLAcAi7IlN","KNWAcAHy5tu","CDFAcAOGWwF","bPeAcA1ptKK","LTHAcAQnUNU","QEOAcAtrH4r","HPXAcA55H8x","bEcAcAcRiNm","OBFAcA0IX6S","bZCAcAReIhd","UWIAcAL8p1V","GWCAcAV8xu7","dVQAcAvqWI6","ONCAcAWWDt1","ZVRAcAN92lS","UZZAcADBdma","BTbAcAfkTNM","CCRAcAHmTXU","USXAcAxz0xY","YdAAcAPVvL8","UNIAcAYk9yf","QUGAcArmWeT","EBLAcAkpPI5","CMRAcACMNDG","YXAAcAWQAip","PBdAcAcDhDh","UICAcAUSAbA","RPYAcAcmxOh","KBEAcArMhNi","YVXAcAY8nJP","XWJAcABihYX","NMJAcAPTHfq","XEfAcANVZdi","TFEAcAxpSGb","URZAcAzRT1l","QaYAcAIyxNh","ODLAcA16nqZ","HLSAcApJgeR","cCIAcAZBGYj","UVAAcA5ARPc","ESAAcAg4UPr","TZQAcAhcJEh","IILAcAFxRiH","UUCAcA9d7ky","EcPAcAuBD0n","WJEAcAXlLeW","eZPAcAYpQcK","FCRAcAFTXXf","KKMAcAhGTHn"]"}
c = Cookie{"server-options-folder_member_over_limit_company_list","["NQHAcAV4DT7","ZdXAcA53D9i","NGZAcAXGt9c","OZQAcALNw5G","PeaAcAvdNVX","MUYAcAEOstW","CXcAcAugGXz","ZSPAcA5pFft","COBAcA9J9uS","FHMAcAL4nBT","JceAcADwK26","XbaAcAr82Xu","SdUAcAUzkyw","XSCAcAaV0Z6","MJDAcAfI6rY","DHZAcAg3w7y","ScIAcAwZ0DZ","VeB9cAKFALF","EaDAcAl3CKD","TXZAcAZ6Xs4","HeUAcAwLAk9","HKaAcA9DfKQ","BUCAcAI30qp","VDfAcAAvcH0","ZGDAcAtRzlJ","YSIAcAARvu1","SSaAcAaMF9I","cFcAcApGEQD","IOO9cAA1JUW","VcfAcA192cI","YbLAcAJ8zEl","WECAcAwYje7","HZKAcAz7NsY","aTKAcAREcWH","dacAcAUbTlI","IJWAcAHAFzZ","fWbAcAyO16U","DEeAcAckoZT","bAQ9cAkaQjr","aEC9cAYenVt","YCbAcAbZr7h","GGAAcAB4aIZ","fVTAcAZ2kvV","DVEAcAuwic7","RAFAcAaANNt","aYNAcA4gKoR","WaQAcAuhMhm","SOdAcAkGnL0","HWEAcAnv6KE","aYfAcAXXRBC","WXcAcAZaqwl","OOAAcAx2j6j","RcYAcANHiZf","WeUAcAH8s5C","DEUAcANq9Yu","dODAcALyZl2","QdbAcAaN6ap","LXaAcAMEuAI","IPEAcA59h5i","BPTAcAk7Ixj","XQaAcAdsj1r","CZJAcA0MlvN","cGOAcA19jBU","FLNAcAO6dWW","fVeAcA9mcfC","PNeAcAc0o2l","OGBAcAUk5lu","GGdAcA5H2df","bZCAcAReIhd","UCUAcA5nQH9","eJSAcA57mZa","fLKAcAL6UnN","OdCAcAAg8OT","BSVAcA1RHQA","ONCAcAWWDt1","JEVAcASAN7b","CTOAcAeJ4YH","USXAcAxz0xY","YdAAcAPVvL8","BGPAcAS5ygS","AEYAcAkdE3E","UFHAcAvn7zu","NJHAcAq7Wc7","EKTAcAh2Nvv","fRTAcAJAJAI","TACAcAKgmRy","cfKAcAZksyx","TFEAcAxpSGb","KASAcA8LgEv"]"}
c = Cookie{"server-options-folders_max_depth","15"}
c = Cookie{"server-options-js_autoreload","false"}
c = Cookie{"server-options-js_autoreload_frequency","1000"}
c = Cookie{"server-options-live_apps_request_pool_timeout","30000"}
c = Cookie{"server-options-long_tasks_metric_threshold","3000"}
c = Cookie{"server-options-max_cell_sections","100000"}
c = Cookie{"server-options-max_company_members","1000"}
c = Cookie{"server-options-max_list_view_children_for_performance","1000"}
c = Cookie{"server-options-max_lotta_cell_sections","1001000"}
c = Cookie{"server-options-max_num_accounts","8"}
c = Cookie{"server-options-max_sections","50000"}
c = Cookie{"server-options-max_undeleted_sections","40000"}
c = Cookie{"server-options-modal_full_text_offscreen_delay","2000"}
c = Cookie{"server-options-modal_full_text_short_query_delay","1000"}
c = Cookie{"server-options-model_check_probe_chance","0.01"}
c = Cookie{"server-options-quip_slack_team_id",""T024JREUW""}
c = Cookie{"server-options-salesforce_query_all_users_limit","10000"}
c = Cookie{"server-options-salesforce_usage_data_start_date","1582670829000"}
c = Cookie{"server-options-salesforce_usage_show_delta_start_ms","1585022447000"}
c = Cookie{"server-options-search_clientperf_sample_fraction","0.02"}
c = Cookie{"server-options-search_index_name",""indextemplate8_minor1""}
c = Cookie{"server-options-sec_before_index_contributions_as_opens","1420099200"}
c = Cookie{"server-options-send_diagnostic_report_for_save_document_errors","false"}
c = Cookie{"server-options-sfdc_email_max_num_accounts","50"}
c = Cookie{"server-options-sfdc_steam_test_types","["test1na44","test1cs2","test1na45","test1cs3","test1na46","test1cs4"]"}
c = Cookie{"server-options-sfdc_steam_types","["stmpa","stmpb"]"}
c = Cookie{"server-options-sfdc_stmf_types","["stmfa","stmfb","stmfc","stmfd"]"}
c = Cookie{"server-options-sfdc_vpod_host",""sfdc_vpod_host""}
c = Cookie{"server-options-show_example_templates","false"}
c = Cookie{"server-options-show_folder_limit_admin_console_alert_banner","false"}
c = Cookie{"server-options-show_folder_limit_admin_console_alert_banner_v2","false"}
c = Cookie{"server-options-sidebar_large_site_threshold","500"}
c = Cookie{"server-options-slides_intro_ga_nux_template_thread_id","null"}
c = Cookie{"server-options-slides_intro_ga_nux_template_thread_secret_path","null"}
c = Cookie{"server-options-staging_server","false"}
c = Cookie{"server-options-throttled_collections_cache_only_threshold","86400000"}
c = Cookie{"server-options-usage_forever_free_companies","["fPQAcAPz5E9","DYTAcAiIFUr","DHZAcAg3w7y","WaQAcAuhMhm","UBdAcAGFnyJ","CPEAcA5Rm2R"]"}
c = Cookie{"server-options-web_browser_id_cookie",""quip_id""}
c = Cookie{"server-options-web_cache_last_updated_cookie",""web_cache_last_updated""}
c = Cookie{"server-options-web_cache_max_staleness_ms","3600000"}
c = Cookie{"session-source-tracking-last-session-id",""463fd830e6ca""}
c = Cookie{"session-source-tracking-last-source-session-id","null"}
c = Cookie{"SLACK_DEBUG_DISABLED","true"}
c = Cookie{"TOGGLE_STORAGE-desktop_channel_versioning_implementation_client_boot","on"}
c = Cookie{"TOGGLE_STORAGE-object_store_redux","treatment"}

	chromedp.Run(ctx, setcookies(
		url,
		cookies,
	))

	time.Sleep(time.Second)
	chromedp.Run(ctx,
		chromedp.Nodes("//*", &res),
	)

	for _, item := range res {

		var innerHTML string
		chromedp.Run(ctx,
			chromedp.InnerHTML(item.FullXPath(), &innerHTML),
		)
		fmt.Println(innerHTML)
	}

	/*
		fmt.Println("1")
		chromedp.Run(ctx,
			chromedp.Navigate(url),
		)
		fmt.Println("2")
		time.Sleep(time.Second)
		sel := `//input[@name='email']`
		chromedp.Run(ctx,
			chromedp.SetValue(sel, email, chromedp.BySearch),
			//chromedp.SendKeys(sel, email),
			chromedp.Click(`//button[@id='submit_btn']`),
		)
		fmt.Println("3")
		//chromedp.Run(ctx,
		//	chromedp.Click(`//button[@id='submit_btn']`),
		//)
		fmt.Println("4")
		//chromedp.Submit(sel).Do(ctx)
		//chromedp.WaitReady("body").Do(ctx)

	*/
}

func setcookies(host string, cookies []Cookie) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			//expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for _, cookie := range cookies {
				err := network.SetCookie(cookie.Name, cookie.Value).
					//WithExpires(&expr).
					WithDomain(cookie.Domain).
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Navigate(host),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := storage.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}
