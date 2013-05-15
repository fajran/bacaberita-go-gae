package feed

import (
	"strings"
	"testing"
)

func getBasicInput() string {
	return `
		<?xml version="1.0" encoding="utf-8"?>
		<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/"
			xmlns:media="http://search.yahoo.com/mrss/"
			xmlns:atom="http://www.w3.org/2005/Atom">
		  <channel>
		    <title>channel title</title>
		    <link>channel link</link>
		    <description>Channel Description</description>
		    <image>
		      <url>Image URL</url>
		      <title>Image Title</title>
		      <link>image link</link>
		    </image>
		    <item>
		      <title>first item title</title>
		      <link>first-item-link</link>
		      <guid isPermaLink="true">guid-of-first-item</guid>
		      <pubDate>Wed, 15 May 2013 15:38:48 +0000</pubDate>
		      <description>one
two
three</description>
			</item>
		    <item>
		      <title>item-2-title</title>
		      <link>
			  item-2-link
			  </link>
		      <guid isPermaLink="true">
			  item-2-guid</guid>
		      <pubDate>Wed, 15 May 2013 03:54:03 +0000</pubDate>
		      <description>item two description</description>
		    </item>
		  </channel>
		</rss>
	`
}

func testBasicInput(rawInput string, t *testing.T) {
	input := []byte(rawInput)

	data, err := Parse(input)
	if err != nil {
		t.Errorf("Unable to parse")
	}
	if data.Title != "channel title" {
		t.Errorf("Invalid title")
	}
	if data.Link != "channel link" {
		t.Errorf("Invalid link")
	}
	if data.Description != "Channel Description" {
		t.Errorf("Invalid description")
	}
	if data.Image.Url != "Image URL" {
		t.Errorf("Invalid image url")
	}
	if data.Image.Title != "Image Title" {
		t.Errorf("Invalid image title")
	}
	if data.Image.Link != "image link" {
		t.Errorf("Invalid image link")
	}
	if len(data.Items) != 2 {
		t.Errorf("Invalid number of items")
	}
}

func TestParseBasic(t *testing.T) {
	rawInput := getBasicInput()
	testBasicInput(rawInput, t)
}

func TestParseBasicOneLine(t *testing.T) {
	rawInput := getBasicInput()
	rawInput = strings.Replace(rawInput, "\n", " ", -1)
	testBasicInput(rawInput, t)
}

func TestParseReddit(t *testing.T) {
	rawInput := `
<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:media="http://search.yahoo.com/mrss/" xmlns:atom="http://www.w3.org/2005/Atom"><channel><title>programming</title><link>http://www.reddit.com/r/programming/</link><description>Computer Programming</description><image><url>http://static.reddit.com/reddit_programming.png</url><title>programming</title><link>http://www.reddit.com/r/programming/</link></image><atom:link rel="self" href="http://www.reddit.com/r/programming/.rss" type="application/rss+xml" /><item><title>New Yorker Launches New Whistleblower Submission System, With Code Written by the Late Aaron Swartz</title><link>http://www.reddit.com/r/programming/comments/1edyl7/new_yorker_launches_new_whistleblower_submission/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edyl7/new_yorker_launches_new_whistleblower_submission/</guid><pubDate>Wed, 15 May 2013 15:38:48 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/gallais&#34;&gt; gallais &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;https://pressfreedomfoundation.org/blog/2013/05/new-yorker-launches-new-whistleblower-submission-system-code-written-late-aaron-swartz&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edyl7/new_yorker_launches_new_whistleblower_submission/"&gt;[10 comments]&lt;/a&gt;</description></item><item><title>So Piranaha Games finally found the cause of their HUD issues in Mechwarrior Online. The explanation of the bug itself is kinda interesting.</title><link>http://www.reddit.com/r/programming/comments/1ed2pv/so_piranaha_games_finally_found_the_cause_of/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ed2pv/so_piranaha_games_finally_found_the_cause_of/</guid><pubDate>Wed, 15 May 2013 03:54:03 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/juju2112&#34;&gt; juju2112 &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://mwomercs.com/forums/topic/117769-hud-bug-brief/&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ed2pv/so_piranaha_games_finally_found_the_cause_of/"&gt;[259 comments]&lt;/a&gt;</description></item><item><title>Google Launches Android Studio And New Features For Developer Console, Including Beta Releases And Staged Rollout</title><link>http://www.reddit.com/r/programming/comments/1ee4vx/google_launches_android_studio_and_new_features/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ee4vx/google_launches_android_studio_and_new_features/</guid><pubDate>Wed, 15 May 2013 17:03:11 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/mtlion&#34;&gt; mtlion &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://techcrunch.com/2013/05/15/google-launches-android-studio-a-development-tool-for-apps/&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ee4vx/google_launches_android_studio_and_new_features/"&gt;[14 comments]&lt;/a&gt;</description></item><item><title>Georgia Tech, in collaboration with UDacity, will offer a fully-accredited, online opencourse Masters Degree program in Computer Science starting Fall 2014 for under $7k</title><link>http://www.reddit.com/r/programming/comments/1ed8qt/georgia_tech_in_collaboration_with_udacity_will/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ed8qt/georgia_tech_in_collaboration_with_udacity_will/</guid><pubDate>Wed, 15 May 2013 05:42:55 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/cybernoodles&#34;&gt; cybernoodles &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.omscs.gatech.edu/&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ed8qt/georgia_tech_in_collaboration_with_udacity_will/"&gt;[150 comments]&lt;/a&gt;</description></item><item><title>Apple’s new Objective-C to Javascript Bridge</title><link>http://www.reddit.com/r/programming/comments/1ee8in/apples_new_objectivec_to_javascript_bridge/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ee8in/apples_new_objectivec_to_javascript_bridge/</guid><pubDate>Wed, 15 May 2013 17:48:09 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/f666&#34;&gt; f666 &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.steamclock.com/blog/2013/05/apple-objective-c-javascript-bridge/&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ee8in/apples_new_objectivec_to_javascript_bridge/"&gt;[13 comments]&lt;/a&gt;</description></item><item><title>Android Studio</title><link>http://www.reddit.com/r/programming/comments/1eeoqu/android_studio/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eeoqu/android_studio/</guid><pubDate>Wed, 15 May 2013 21:06:29 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/linucs&#34;&gt; linucs &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://developer.android.com/sdk/installing/studio.html&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eeoqu/android_studio/"&gt;[comment]&lt;/a&gt;</description></item><item><title>How Estonia became E-stonia</title><link>http://www.reddit.com/r/programming/comments/1edixy/how_estonia_became_estonia/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edixy/how_estonia_became_estonia/</guid><pubDate>Wed, 15 May 2013 10:39:16 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/yrmjy&#34;&gt; yrmjy &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.bbc.co.uk/news/business-22317297&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edixy/how_estonia_became_estonia/"&gt;[29 comments]&lt;/a&gt;</description></item><item><title>Compiling a C dialect straight to ELF64 with a single small OCaml file</title><link>http://www.reddit.com/r/programming/comments/1edvqj/compiling_a_c_dialect_straight_to_elf64_with_a/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edvqj/compiling_a_c_dialect_straight_to_elf64_with_a/</guid><pubDate>Wed, 15 May 2013 14:57:23 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/_mpu&#34;&gt; _mpu &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://c9x.me/qcc&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edvqj/compiling_a_c_dialect_straight_to_elf64_with_a/"&gt;[16 comments]&lt;/a&gt;</description></item><item><title>Parallelism and concurrency need different tools</title><link>http://www.reddit.com/r/programming/comments/1edvqd/parallelism_and_concurrency_need_different_tools/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edvqd/parallelism_and_concurrency_need_different_tools/</guid><pubDate>Wed, 15 May 2013 14:57:20 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/pdq&#34;&gt; pdq &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.yosefk.com/blog/parallelism-and-concurrency-need-different-tools.html&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edvqd/parallelism_and_concurrency_need_different_tools/"&gt;[2 comments]&lt;/a&gt;</description></item><item><title>An Interview with Bjarne Stroustrup | InformIT // cross post from /r/cpp</title><link>http://www.reddit.com/r/programming/comments/1ee9gh/an_interview_with_bjarne_stroustrup_informit/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ee9gh/an_interview_with_bjarne_stroustrup_informit/</guid><pubDate>Wed, 15 May 2013 18:00:30 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/mttd&#34;&gt; mttd &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.informit.com/articles/article.aspx?p=2080042&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ee9gh/an_interview_with_bjarne_stroustrup_informit/"&gt;[9 comments]&lt;/a&gt;</description></item><item><title>Google's new AppEngine language is PHP</title><link>http://www.reddit.com/r/programming/comments/1eekcz/googles_new_appengine_language_is_php/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eekcz/googles_new_appengine_language_is_php/</guid><pubDate>Wed, 15 May 2013 20:14:46 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/jiunec&#34;&gt; jiunec &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_PHP&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eekcz/googles_new_appengine_language_is_php/"&gt;[3 comments]&lt;/a&gt;</description></item><item><title>Where Is .NET Headed?</title><link>http://www.reddit.com/r/programming/comments/1edtfn/where_is_net_headed/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edtfn/where_is_net_headed/</guid><pubDate>Wed, 15 May 2013 14:24:08 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/stusmith&#34;&gt; stusmith &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://odetocode.com/blogs/scott/archive/2013/05/15/where-is-net-headed.aspx&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edtfn/where_is_net_headed/"&gt;[4 comments]&lt;/a&gt;</description></item><item><title>Learning Programming is Not a Race</title><link>http://www.reddit.com/r/programming/comments/1edshm/learning_programming_is_not_a_race/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edshm/learning_programming_is_not_a_race/</guid><pubDate>Wed, 15 May 2013 14:09:22 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/emmett9001&#34;&gt; emmett9001 &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://blog.parsely.com/post/50495864825/to-the-next-parse-ly-intern-learning-is-not-a-race&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edshm/learning_programming_is_not_a_race/"&gt;[7 comments]&lt;/a&gt;</description></item><item><title>DConf 2013 Day 1 Talk 4: Writing Testable Code in D by Ben Gertzfield</title><link>http://www.reddit.com/r/programming/comments/1edih2/dconf_2013_day_1_talk_4_writing_testable_code_in/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edih2/dconf_2013_day_1_talk_4_writing_testable_code_in/</guid><pubDate>Wed, 15 May 2013 10:23:06 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/andralex&#34;&gt; andralex &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://youtube.com/watch?v=V98Z11V7kEY&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edih2/dconf_2013_day_1_talk_4_writing_testable_code_in/"&gt;[7 comments]&lt;/a&gt;</description></item><item><title>Terra – a low-level counterpart to Lua</title><link>http://www.reddit.com/r/programming/comments/1ebow2/terra_a_lowlevel_counterpart_to_lua/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ebow2/terra_a_lowlevel_counterpart_to_lua/</guid><pubDate>Tue, 14 May 2013 17:05:49 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/burntsushi&#34;&gt; burntsushi &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://terralang.org&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ebow2/terra_a_lowlevel_counterpart_to_lua/"&gt;[145 comments]&lt;/a&gt;</description></item><item><title>Bitbucket introduces online editing</title><link>http://www.reddit.com/r/programming/comments/1eccsg/bitbucket_introduces_online_editing/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eccsg/bitbucket_introduces_online_editing/</guid><pubDate>Tue, 14 May 2013 22:03:04 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/jespern&#34;&gt; jespern &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://blog.bitbucket.org/2013/05/14/edit-your-code-in-the-cloud-with-bitbucket/&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eccsg/bitbucket_introduces_online_editing/"&gt;[69 comments]&lt;/a&gt;</description></item><item><title>Compiler internals: implementation of RTTI and exceptions in MSVC and GCC [PDF]</title><link>http://www.reddit.com/r/programming/comments/1edlyi/compiler_internals_implementation_of_rtti_and/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1edlyi/compiler_internals_implementation_of_rtti_and/</guid><pubDate>Wed, 15 May 2013 12:03:13 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/igor_sk&#34;&gt; igor_sk &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.hexblog.com/wp-content/uploads/2012/06/Recon-2012-Skochinsky-Compiler-Internals.pdf&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1edlyi/compiler_internals_implementation_of_rtti_and/"&gt;[1 comment]&lt;/a&gt;</description></item><item><title>Amazedsaint's Tech Journal: May be this is the best time for Microsoft to Open Source the .NET Framework</title><link>http://www.reddit.com/r/programming/comments/1eeby4/amazedsaints_tech_journal_may_be_this_is_the_best/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eeby4/amazedsaints_tech_journal_may_be_this_is_the_best/</guid><pubDate>Wed, 15 May 2013 18:30:24 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/1nvader&#34;&gt; 1nvader &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.amazedsaint.com/2013/05/may-be-this-is-best-time-for-microsoft.html&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eeby4/amazedsaints_tech_journal_may_be_this_is_the_best/"&gt;[3 comments]&lt;/a&gt;</description></item><item><title>Linux code is the 'benchmark of quality,' study concludes</title><link>http://www.reddit.com/r/programming/comments/1ee2dj/linux_code_is_the_benchmark_of_quality_study/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ee2dj/linux_code_is_the_benchmark_of_quality_study/</guid><pubDate>Wed, 15 May 2013 16:32:01 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/Campers&#34;&gt; Campers &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.pcworld.in/news/linux-code-benchmark-quality-study-concludes-98752013&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ee2dj/linux_code_is_the_benchmark_of_quality_study/"&gt;[2 comments]&lt;/a&gt;</description></item><item><title>Drawing Dynamic Visualizations</title><link>http://www.reddit.com/r/programming/comments/1ed7wh/drawing_dynamic_visualizations/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ed7wh/drawing_dynamic_visualizations/</guid><pubDate>Wed, 15 May 2013 05:24:54 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/nexes300&#34;&gt; nexes300 &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://vimeo.com/66085662&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ed7wh/drawing_dynamic_visualizations/"&gt;[5 comments]&lt;/a&gt;</description></item><item><title>The Elusive Universal Web ByteCode</title><link>http://www.reddit.com/r/programming/comments/1ecwvg/the_elusive_universal_web_bytecode/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ecwvg/the_elusive_universal_web_bytecode/</guid><pubDate>Wed, 15 May 2013 02:32:02 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/therapy&#34;&gt; therapy &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://mozakai.blogspot.com/2013/05/the-elusive-universal-web-bytecode.html&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ecwvg/the_elusive_universal_web_bytecode/"&gt;[50 comments]&lt;/a&gt;</description></item><item><title>A week of no media consumption</title><link>http://www.reddit.com/r/programming/comments/1eerw0/a_week_of_no_media_consumption/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eerw0/a_week_of_no_media_consumption/</guid><pubDate>Wed, 15 May 2013 21:44:47 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/evjan&#34;&gt; evjan &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://peterevjan.com/posts/a-week-of-no-media-consumption/&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eerw0/a_week_of_no_media_consumption/"&gt;[comment]&lt;/a&gt;</description></item><item><title>As a Christian software engineer, I enjoyed this. The parable of the two computer programs</title><link>http://www.reddit.com/r/programming/comments/1eerjr/as_a_christian_software_engineer_i_enjoyed_this/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eerjr/as_a_christian_software_engineer_i_enjoyed_this/</guid><pubDate>Wed, 15 May 2013 21:40:25 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/devroth&#34;&gt; devroth &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://creation.com/parable-two-computer-programs&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eerjr/as_a_christian_software_engineer_i_enjoyed_this/"&gt;[comment]&lt;/a&gt;</description></item><item><title>Binding D To C</title><link>http://www.reddit.com/r/programming/comments/1eerd9/binding_d_to_c/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1eerd9/binding_d_to_c/</guid><pubDate>Wed, 15 May 2013 21:38:08 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/tanczosm&#34;&gt; tanczosm &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://www.gamedev.net/page/resources/_/technical/game-programming/binding-d-to-c-r3122&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1eerd9/binding_d_to_c/"&gt;[comment]&lt;/a&gt;</description></item><item><title>Abstractivate: Two Models of Computation: or, Why I'm Switching Trains</title><link>http://www.reddit.com/r/programming/comments/1ecyko/abstractivate_two_models_of_computation_or_why_im/</link><guid isPermaLink="true">http://www.reddit.com/r/programming/comments/1ecyko/abstractivate_two_models_of_computation_or_why_im/</guid><pubDate>Wed, 15 May 2013 02:54:52 +0000</pubDate><description>submitted by &lt;a href=&#34;http://www.reddit.com/user/yogthos&#34;&gt; yogthos &lt;/a&gt; &lt;br/&gt; &lt;a href=&#34;http://blog.jessitron.com/2013/05/two-models-of-computation-or-why-im.html&#34;&gt;[link]&lt;/a&gt; &lt;a href="http://www.reddit.com/r/programming/comments/1ecyko/abstractivate_two_models_of_computation_or_why_im/"&gt;[60 comments]&lt;/a&gt;</description></item></channel></rss>
`
	input := []byte(rawInput)

	data, err := Parse(input)
	if err != nil {
		t.Errorf("Unable to parse")
	}
	if data.Title != "programming" {
		t.Errorf("Invalid title")
	}
	if data.Link != "http://www.reddit.com/r/programming/" {
		t.Errorf("Invalid link")
	}
	if data.Description != "Computer Programming" {
		t.Errorf("Invalid description")
	}
	if data.Image.Url != "http://static.reddit.com/reddit_programming.png" {
		t.Errorf("Invalid image url")
	}
	if data.Image.Title != "programming" {
		t.Errorf("Invalid image title")
	}
	if data.Image.Link != "http://www.reddit.com/r/programming/" {
		t.Errorf("Invalid image link")
	}
	if len(data.Items) != 25 {
		t.Errorf("Invalid number of items")
	}
}
