Go package for 'https://github.com/valeriangalliat/spotify-buddylist'

Why?
    - Spotify's official API does not provide support for viewing the activity of people you follow, this package implements it with a few of Spotify's internal APIs


Parameters needed:
    - Cookie: 
        - After you log into the Spotify web player, you need to find the cookie provided named 'sp_dc' and get the value of it
        - This cookie expires every year, so if used in production, it will be best to automate the collection of it

