ignore /stylesheets\/((?!main)(?!typography)).*\.css.scss/

configure :build do
  activate :minify_css
  activate :minify_javascript
end

activate :blog do |blog|
  blog.name = "blog"
  blog.prefix = "blog"
  blog.permalink = "{title}"
  blog.generate_day_pages = false
  blog.generate_month_pages = false
  blog.generate_tag_pages = false
  blog.generate_year_pages = false
  blog.layout = "post"
  blog.sources = "{id}-{title}.html"
  blog.default_extension = "md"
end

activate :blog do |blog|
  blog.name = "linklog"
  blog.prefix = "linklog"
  blog.permalink = "{title}"
  blog.generate_day_pages = false
  blog.generate_month_pages = false
  blog.generate_tag_pages = false
  blog.generate_year_pages = false
  blog.layout = "post"
  blog.sources = "{year}-{month}-{day}-{title}.html"
  blog.default_extension = "md"
end

activate :blog do |blog|
  blog.name = "birch-devlog"
  blog.prefix = "birch-devlog"
  blog.permalink = "{title}"
  blog.generate_day_pages = false
  blog.generate_month_pages = false
  blog.generate_tag_pages = false
  blog.generate_year_pages = false
  blog.layout = "post"
  blog.sources = "{year}-{month}-{day}-{title}.html"
  blog.default_extension = "md"
end

activate :directory_indexes
