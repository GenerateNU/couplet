// Generates dummy test data
package main

import (
	"context"
	"couplet/internal/api"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"slices"

	"github.com/google/uuid"
)

var (
	orgImages = []url.URL{mustParseUrl("https://static01.nyt.com/images/2006/12/07/arts/08ica600.1.jpg"),
		mustParseUrl("https://media.architecturaldigest.com/photos/585c57b19a1af9cb3992ee41/1:1/w_3754,h_3754,c_limit/beaux-arts-paris-06.jpg"),
		mustParseUrl("https://ids.si.edu/ids/deliveryService?id=https://www.si.edu/sites/default/files/newsdesk/building/aib-03_print.jpg&max_w=600"),
		mustParseUrl("https://images.adsttc.com/media/images/5ffe/5a97/63c0/174c/f800/00ee/newsletter/1.jpg?1610504845")}
	orgTags        = []string{"nonprofit", "family-owned", "international", "museum", "university", "eco-friendly", "start-up"}
	eventAddresses = []api.Address{{Line1: "360 Huntington Ave", ZipCode: "02115", City: "Boston", State: "MA", Country: "US"},
		{Line1: "465 Huntington Ave", ZipCode: "02115", City: "Boston", State: "MA", Country: "US"},
		{Line1: "25 Harbor Shore Dr", ZipCode: "02210", City: "Boston", State: "MA", Country: "US"}}
	eventImages = []url.URL{mustParseUrl("https://d1nn9x4fgzyvn4.cloudfront.net/styles/scaled_562_wide/s3/2023-08/0289_4x3.jpg?itok=g1xziFrq"),
		mustParseUrl("https://umanitoba.ca/art/sites/art/files/styles/21x9_1100w/public/2020-08/exhibitions-events.jpg?itok=ih_87Wlz"),
		mustParseUrl("https://www.freemanarts.org/de/cache/content/30/hr_Events_Tickets_Hero_2022.png"),
		mustParseUrl("https://365thingsinhouston.com/wp-content/uploads/2024/01/top-things-to-do-this-week-in-houston-january-1-7-2024-tina-turner-musical-2.jpg"),
		mustParseUrl("https://ncartmuseum.org/wp-content/uploads/elementor/thumbs/DSC05129-scaled-qjgna9937kddrhq77qvjzryn2g4etat2ucgn4j77r2.jpg"),
		mustParseUrl("https://bostonchildrensmuseum.org/wp-content/uploads/2022/03/CRW_5564-1.jpg"),
		mustParseUrl("https://www.metmuseum.org/-/media/images/join-and-give/host-an-event/host-an-event_block.jpg?sc_lang=en"),
		mustParseUrl("https://www.cambridgema.gov/-/media/Images/publicworks/specialevents/Danehy/DanehyFamilyDay_KyleKlein_KKP17278.jpg?mw=1920"),
		mustParseUrl("https://www.bendparksandrec.org/wp-content/uploads/2017/12/Riverbend-Park-Community-Event-Rentals.jpg"),
		mustParseUrl("https://lastatehistoricpark.org/wp-content/uploads/2019/05/Website-Facebook-Share-Yoast-SEO-Picture-1200x630.png"),
		mustParseUrl("https://img.mlbstatic.com/mlb-images/image/upload/t_5x2/t_w2208/mlb/a1bcnwaokjzl25d37u5v.jpg"),
		mustParseUrl("https://gggp.org/wp-content/uploads/2023/11/eventhero_flowerpiano.jpg"),
		mustParseUrl("https://citytableboston.com/wp-content/uploads/2019/07/www.citytableboston.com685city-table-back-bay-holiday-parties-private-room.png"),
		mustParseUrl("https://www.upmenu.com/wp-content/uploads/2021/07/3-restaurant-event-ideas-example-food-tastings.jpg"),
		mustParseUrl("https://www.buzztime.com/business/wp-content/uploads/2019/08/shutterstock_365582531.jpg")}
	eventExternalLink = mustParseUrl("https://www.google.com/")
	eventTags         = []string{"indoors", "outdoors", "art", "music", "food", "active", "limited-time", "showcase", "performance"}
	userImages        = []url.URL{mustParseUrl("https://static01.nyt.com/images/2015/08/10/fashion/10TELLER1/10TELLER1-superJumbo.jpg"),
		mustParseUrl("https://hollywoodlife.com/wp-content/uploads/2015/10/terry-crews-bio-photo.jpg?quality=100"),
		mustParseUrl("https://www.nydailynews.com/wp-content/uploads/migration/2010/02/04/T2RCYCF37IKCK7R544GN6KQUEE.jpg?w=535"),
		mustParseUrl("https://assets.vogue.com/photos/65418f5726fcdeb5a090adf8/master/w_2560%2Cc_limit/1530545900"),
		mustParseUrl("https://img.buzzfeed.com/buzzfeed-static/static/2020-11/12/21/asset/b537877860a1/sub-buzz-2914-1605214930-12.jpg?downsize=700%3A%2A&output-quality=auto&output-format=auto"),
		mustParseUrl("https://hips.hearstapps.com/hmg-prod/images/766/shutterstock-350127209-1515591195.jpg?resize=640:*"),
		mustParseUrl("https://coveredgeekly.com/wp-content/uploads/Top-15-Most-Beautiful-Female-Celebrities-Actresses-of-2023-According-to-Polls-Image-3-1024x549.jpg"),
		mustParseUrl("https://www.southernliving.com/thmb/lPaazTFUvGagaO5nNducVCX8j8M=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/jennifer-lawerence-529843874-1-3d7e3619a9e843c48590b7fa22f55bdf.jpg"),
		mustParseUrl("https://footwearnews.com/wp-content/uploads/2019/10/cardi-b-two-tone-suit.jpg?w=600&h=337&crop=1"),
		mustParseUrl("https://imgix.ranker.com/list_img_v2/7162/2727162/original/most-beautiful-female-celebrities-2018?fit=crop&fm=pjpg&q=80&dpr=2&w=1200&h=720"),
		mustParseUrl("https://hips.hearstapps.com/hmg-prod/images/taylor-swift-1675879804.png?crop=0.285xw:0.570xh;0.673xw,0&resize=640:*"),
		mustParseUrl("https://imgix.ranker.com/user_node_img/3707/74120030/original/74120030-photo-u-1084707462"),
		mustParseUrl("https://wl-brightside.cf.tsp.li/resize/728x/jpg/061/cb6/6da4fd5a66b816c102013032e2.jpg"),
		mustParseUrl("https://media.glamour.com/photos/5dc6fbd07b1dcc0008dc200a/master/pass/GettyImages-1186240491.jpg"),
		mustParseUrl("https://assets.teenvogue.com/photos/569699c72a465c9c5e41b564/1:1/w_2976,h_2976,c_limit/143479663.jpg"),
		mustParseUrl("https://a57.foxnews.com/static.foxnews.com/foxnews.com/content/uploads/2018/12/1200/675/friends-cast-getty.jpg?ve=1&tl=1"),
		mustParseUrl("https://variety.com/wp-content/uploads/2020/02/friends.jpg?w=1000"),
		mustParseUrl("https://digitalspyuk.cdnds.net/17/09/2560x1704/1488375386-1488370326-friends-cast.jpg"),
		mustParseUrl("https://imgix.ranker.com/list_img_v2/6002/2806002/original/what-the-office-cast-thinks-of-show?fit=crop&fm=pjpg&q=80&dpr=2&w=1200&h=720"),
		mustParseUrl("https://roost.nbcuni.com/bin/viewasset.html/content/dam/Peacock/Landing-Pages/2-0-design/the-office/cast-the-office-dwight-schrute.jpg/_jcr_content/renditions/original.JPEG"),
		mustParseUrl("https://www.usmagazine.com/wp-content/uploads/2020/05/That-70s-Show-Cast-Where-Are-They-Now-Feature.jpg?quality=82&strip=all"),
		mustParseUrl("https://www.digitaltrends.com/wp-content/uploads/2022/04/that-70s-show-tv.jpg?p=1"),
		mustParseUrl("https://grantland.com/wp-content/uploads/2013/08/grant_fox_70sshow_64011.jpg?w=750"),
		mustParseUrl("https://ew.com/thmb/cutRxd7OFNPWVv3KMVXrg3ReeJw=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/seinfeld-1-f70fc02a753e416dafc99bba7a1e76ad.jpg"),
		mustParseUrl("https://www.hollywoodreporter.com/wp-content/uploads/2018/05/cast_37310_copy_-_h_2018.jpg"),
		mustParseUrl("https://www.womansworld.com/wp-content/uploads/2024/01/cast-of-seinfeld.jpg"),
		mustParseUrl("https://static1.colliderimages.com/wordpress/wp-content/uploads/2023/09/seinfeld-1.jpg"),
		mustParseUrl("https://preview.redd.it/in-your-opinion-which-episode-is-peak-george-v0-aoopuvla4gy91.jpg?auto=webp&s=4440c9d751215a23069ec87c2d021c3ca298440a"),
		mustParseUrl("https://press.hulu.com/app/uploads/NewGirl-Season792-copy.png"),
		mustParseUrl("https://images.immediate.co.uk/production/volatile/sites/3/2021/05/keyart-09cc4c7.jpg?quality=90&resize=620,414"),
		mustParseUrl("https://i.pinimg.com/736x/67/35/4b/67354b6d373b959448a4543218085a7c.jpg"),
		mustParseUrl("https://bgr.com/wp-content/uploads/2023/03/New-Girl.jpg?quality=82&strip=all&resize=1400,1400"),
		mustParseUrl("https://i.insider.com/629a474d7bc6a80018b63cd0?width=800&format=jpeg&auto=webp")}
)

func main() {
	// Gather flags
	port := flag.Uint("port", 8080, "the server port to connect to")
	numOrgs := flag.Uint("orgs", 0, "the number of orgs to generate")
	numEvents := flag.Uint("events", 0, "the number of events to generate")
	numUsers := flag.Uint("users", 0, "the number of users to generate")
	numEventSwipes := flag.Uint("eventSwipes", 0, "the number of event swipes to generate")
	numUserSwipes := flag.Uint("userSwipes", 0, "the number of user swipes to generate")
	flag.Parse()

	// Create client
	ctx := context.Background()
	client, err := api.NewClient(fmt.Sprintf("http://localhost:%d", *port))
	if err != nil {
		log.Fatalln(err)
	}

	// Generate orgs
	fmt.Printf("generating %d org(s)...\n", *numOrgs)
	orgIds := []uuid.UUID{}
	for i := uint(0); i < *numOrgs; i++ {
		// Define org
		newOrg := api.OrgsPostReq{}
		newOrg.Name = fmt.Sprintf("org-%d", i)
		newOrg.Bio = fmt.Sprintf("At %s, we connect people through events", newOrg.Name)
		newOrg.Image = api.NewOptURI(orgImages[rand.Intn(len(orgImages))])
		newOrg.Tags = []string{}
		for j := 0; j < rand.Intn(5); j++ {
			tag := orgTags[rand.Intn(len(orgTags))]
			if !slices.Contains(newOrg.Tags, tag) {
				newOrg.Tags = append(newOrg.Tags, tag)
			}
		}

		// Create org
		res, err := client.OrgsPost(ctx, &newOrg)
		if err != nil {
			log.Fatalln(err)
		}
		resCreated, ok := res.(*api.OrgsPostCreated)
		if !ok {
			log.Fatalln("failed to create org")
		}
		orgIds = append(orgIds, resCreated.ID)
	}
	fmt.Printf("\tgenerated %d org(s)\n", len(orgIds))

	// Generate events
	fmt.Printf("generating %d event(s)...\n", *numEvents)
	eventIds := []uuid.UUID{}
	for i := uint(0); i < *numEvents; i++ {
		// Define event
		newEvent := api.EventsPostReq{}
		newEvent.Name = fmt.Sprintf("event-%d", i)
		newEvent.Bio = fmt.Sprintf("Come to %s and have the best night of your life!", newEvent.Name)
		newEvent.Address = api.NewOptAddress(eventAddresses[rand.Intn(len(eventAddresses))])
		newEvent.Images = []url.URL{}
		for j := 0; j < 1+rand.Intn(4); j++ {
			image := eventImages[rand.Intn(len(eventImages))]
			if !slices.Contains(newEvent.Images, image) {
				newEvent.Images = append(newEvent.Images, image)
			}
		}
		if rand.Intn(2) == 0 {
			newEvent.ExternalLink = api.NewOptURI(eventExternalLink)
		}
		newEvent.Tags = []string{}
		for j := 0; j < rand.Intn(5); j++ {
			tag := eventTags[rand.Intn(len(eventTags))]
			if !slices.Contains(newEvent.Tags, tag) {
				newEvent.Tags = append(newEvent.Tags, tag)
			}
		}
		newEvent.OrgId = orgIds[rand.Intn(len(orgIds))]

		// Create event
		res, err := client.EventsPost(ctx, &newEvent)
		if err != nil {
			log.Fatalln(err)
		}
		resCreated, ok := res.(*api.EventsPostCreated)
		if !ok {
			log.Fatalln("failed to create event")
		}
		eventIds = append(eventIds, resCreated.ID)
	}
	fmt.Printf("\tgenerated %d event(s)\n", len(eventIds))

	// Generate users
	fmt.Printf("generating %d user(s)...\n", *numUsers)
	userIds := []uuid.UUID{}
	for i := uint(0); i < *numUsers; i++ {
		// Define user
		newUser := api.UsersPostReq{}
		newUser.FirstName = fmt.Sprintf("user-%d", i)
		newUser.LastName = "lastname"
		newUser.Age = uint8(18 + rand.Intn(23))
		newUser.Images = []url.URL{}
		for j := 0; j < 1+rand.Intn(4); j++ {
			image := userImages[rand.Intn(len(userImages))]
			if !slices.Contains(newUser.Images, image) {
				newUser.Images = append(newUser.Images, image)
			}
		}

		// Create user
		res, err := client.UsersPost(ctx, &newUser)
		if err != nil {
			log.Fatalln(err)
		}
		resCreated, ok := res.(*api.UsersPostCreated)
		if !ok {
			log.Fatalln("failed to create user")
		}
		userIds = append(userIds, resCreated.ID)
	}
	fmt.Printf("\tgenerated %d user(s)\n", len(userIds))

	// Generate event swipes
	fmt.Printf("generating %d event swipe(s)...\n", *numEventSwipes)
	eventSwipes := []struct {
		user  uuid.UUID
		event uuid.UUID
	}{}
	for i := uint(0); i < *numEventSwipes; i++ {
		// Define event swipe
		newEventSwipe := api.EventSwipe{}
		newEventSwipe.UserId = userIds[rand.Intn(len(userIds))]
		newEventSwipe.EventId = eventIds[rand.Intn(len(eventIds))]
		if rand.Intn(2) == 1 {
			newEventSwipe.Liked = true
		}
		pair := struct {
			user  uuid.UUID
			event uuid.UUID
		}{user: newEventSwipe.UserId, event: newEventSwipe.EventId}
		if slices.Contains(eventSwipes, pair) {
			fmt.Println("\tevent swipe already exists, skipping...")
			continue
		}

		// Create event swipe
		res, err := client.EventsSwipesPost(ctx, &newEventSwipe)
		if err != nil {
			log.Fatalln(err)
		}
		_, ok := res.(*api.EventSwipe)
		if !ok {
			log.Fatalln("failed to create event swipe")
		}
		eventSwipes = append(eventSwipes, pair)
	}
	fmt.Printf("\tgenerated %d event swipe(s)\n", len(eventSwipes))

	// Generate user swipes
	fmt.Printf("generating %d user swipe(s)...\n", *numUserSwipes)
	userSwipes := []struct {
		user  uuid.UUID
		other uuid.UUID
	}{}
	for i := uint(0); i < *numUserSwipes; i++ {
		// Define user swipe
		newUserSwipe := api.UserSwipe{}
		newUserSwipe.UserId = userIds[rand.Intn(len(userIds))]
		for (newUserSwipe.OtherUserId == uuid.UUID{} || newUserSwipe.UserId == newUserSwipe.OtherUserId) {
			newUserSwipe.OtherUserId = userIds[rand.Intn(len(userIds))]
		}
		pair := struct {
			user  uuid.UUID
			other uuid.UUID
		}{user: newUserSwipe.UserId, other: newUserSwipe.OtherUserId}
		if slices.Contains(userSwipes, pair) {
			fmt.Println("\tuser swipe already exists, skipping...")
			continue
		}

		if rand.Intn(2) == 1 {
			newUserSwipe.Liked = true
		}

		// Create user swipe
		res, err := client.UsersSwipesPost(ctx, &newUserSwipe)
		if err != nil {
			log.Fatalln(err)
		}
		_, ok := res.(*api.UserSwipe)
		if !ok {
			log.Fatalln("failed to create user swipe")
		}
		userSwipes = append(userSwipes, pair)
	}
	fmt.Printf("\tgenerated %d user swipe(s)\n", len(userSwipes))
}

// Parses a string as a URL. Panics if parsing fails
func mustParseUrl(rawURL string) url.URL {
	url, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return *url
}
