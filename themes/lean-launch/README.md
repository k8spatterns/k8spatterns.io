# Hugo Theme: Lean Launch Page

This theme is to build a quick, single page launch site. You may want to build a site like this when you have a business/start up/product idea that you'd like to validate. This is a common practice nowadays by people following The Lean Startup Methodology.

This theme allows you to announce a company or product, provide social media links, and a email form for users to signup which is hooked up to MailChimp.

## Getting Started

An example of how to use this theme is provided in the `./exampleSite/` directory. Basically, you'll want to use your Hugo config file as usual and customize all of the variables in the `params` section. The example Hugo config can be found at `./exampleSite/config.toml`.

### Google Analytics

You can provide your Google Analytics ID in the Hugo config to add GA tracking to your page. To measure your launch page's effectiveness, this is highly suggested.

### Social Media

You can set the social media URL for Twitter, Facebook, Instagram, and Google+ to your Hugo config file and the associated icons (and links) will show up on your launch page.

### MailChimp

Set `mailchimpURL` in your Hugo config to enable an email form on your page allowing you to collect signups from visitors. This how is how you build your launch list for your product.

#### Getting the URL

To get the MailChimp URL, go to your list on MailChimp, click "Signup forms", and select "Embedded forms". In the "Copy/paste onto your site" text area, you'll see HTMl that says `<form action="`. The URL that follows is what you want.

## Contribute

This theme is under active development. Found a bug? Have feedback? Let us know on GitHub.
