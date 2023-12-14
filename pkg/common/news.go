package common

import "errors"

type NewsInfo struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	ImageID string `json:"image_id"`
}
type News struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	ImageID string `json:"image_id"`
}

var newsList = []News{
	{
		ID:      1,
		Title:   "A Culinary Carnival on TWP!",
		Content: "🍃 Boost Your Health: Try Our New White Powders! 🍃\n\nWe're excited to announce the launch of our new line of white powders! Sourced from the finest organic farms, our powders are packed with nutrients to supercharge your day. Explore options like Spirulina, Chlorella, and Wheatgrass. Perfect for smoothies, shakes, and baking. Start enhancing your meals with a powerhouse of nutrition today!",
		ImageID: "https://media.discordapp.net/attachments/905456405597802546/1184817363032490095/file-QKA0h3P3YmJHXJ3tZTraBuyd.png",
	},
	{
		ID:      2,
		Title:   "Enter the Realm of the New Year!",
		Content: "📅 As we eagerly flip towards the New Year, Los Pollos Hermanos is thrilled to bring a gastronomic revolution exclusively on TWP. 🎉 In this festive season, we've selected an array of our most exquisite Powders, each brimming with the joy and vibrancy of the New Year. 🌟 We're delighted to invite you to join our celebration with an irresistible offer - a whopping 100% off our beloved collection! 🎈\nEmbark on a flavor-packed journey as you explore our labyrinth of tastes, each more captivating than the last. 🍲 Our Powders are more than just seasonings; they're carriers of age-old traditions, bringing the essence of distant lands right to your kitchen. This New Year, let your meals be a canvas adorned with the bright hues and rich textures of our premium spices, turning everyday dishes into culinary wonders. 🌍\n\n🎊 Revel in the New Year's festivities with a world of flavors at your fingertips. \n🥘 Create gastronomic masterpieces with the finest ingredients, handpicked by Los Pollos Hermanos.\n🎁 Gift your loved ones the essence of global cuisine, a treasure chest of spices from across the globe.\n\nThis festive season, let not just lanterns light up your home but also the tantalizing aromas from your kitchen, thanks to Los Pollos Hermanos's special New Year offer. 🏮 We believe the right spice isn't just a part of a recipe, but a gateway to an experience that unites friends and family, creating unforgettable moments over shared meals and laughter. 🥂\n\nVisit Los Pollos Hermanos's store on TWP, where each spice jar tells a story, ready to add its magic to your next culinary adventure. 🌟 Here's to a New Year filled with as much flavor as the feasts you'll craft with our finest selections. May the joy of the New Year fill your home with warmth and your heart with happiness. 🏡❤️",
		ImageID: "https://cdn.discordapp.com/attachments/905456405597802546/1184811425747705886/exterior_interactivelightmixcopy.png",
	},
	{
		ID:      3,
		Title:   "Last Chance to get your hands on the best!",
		Content: "🚀 Are you ready to experience the pinnacle of quality in powder products? Look no further! Our C2C (Customer-to-Customer) powder selling website is offering an unmissable opportunity. This is your LAST CHANCE to dive into a world of premium, carefully curated powders at unbeatable prices. \n🌿 Whether you're a fitness enthusiast looking for that perfect protein powder, a beauty aficionado in search of natural skincare solutions, or a culinary expert seeking exotic spices, we've got you covered. Our diverse range of powders caters to all your needs. \n🛒 Shopping with us is not just a transaction; it's an experience. Enjoy the ease of connecting directly with sellers, ensuring you get authentic and high-quality products every time. Plus, with user ratings and reviews, you can shop with confidence. \n💸 Don't miss out on our exclusive deals and discounts. This is your chance to stock up on your favorites or explore new finds. But hurry, these offers won't last forever! \n🔥 Act fast and be part of a community that values quality and variety. Visit our website now and join the ranks of satisfied customers who've discovered the best in powder products. Remember, this is your LAST CHANCE to grab these deals before they're gone for good! \n👉 Visit us at [Your Website Link] and elevate your shopping experience today! \n",
		ImageID: "https://media.discordapp.net/attachments/905456405597802546/1184815114466115635/file-n44FSn494CNNFFTvCAgCuL8o.png",
	},
}

func GetNewsInfo() []NewsInfo {
	var newsInfoList []NewsInfo
	for _, news := range newsList {
		newsInfoList = append(newsInfoList, NewsInfo{
			ID:      news.ID,
			Title:   news.Title,
			ImageID: news.ImageID,
		})
	}
	return newsInfoList
}

func GetNews(id int32) (*News, error) {
	for _, news := range newsList {
		if news.ID == id {
			return &news, nil
		}
	}
	return nil, errors.New("news not found")
}
