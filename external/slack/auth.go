package slack

/*
curl -X GET \
  'https://slack.com/oauth/v2/authorize?client_id=<CLIENT_ID>&scope=<SCOPE>&redirect_uri=<REDIRECT_URI>'

	'https://slack.com/oauth/v2/authorize?client_id=5544062518454.5574454550576&user_scope=chat:write,channels:history,channels:read,channels:write,users:read,users:write&redirect_uri=https://many.pw/slack/callback'


	curl -X POST \
  https://slack.com/api/oauth.v2.access \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'client_id=5544062518454.5574454550576&client_secret=x&code=a&redirect_uri=https://many.pw/slack'

*/
