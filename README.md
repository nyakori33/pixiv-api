# pixiv-api
Golang package for Pixiv Web API.

# API List
## Completed
* GetUserProfile (struct only)
* GetUserProfileSimple (struct only)

## Coming soon
* GetUserHomeArtwork
* GetUserAllArtwork
* GetUserLatestArtwork
* GetArtworkProfile
* GetArtworkTags
* GetArtworkURL
* GetArtworkComments
* GetArtworkRecommend
* GetUgoria[^1]Profile
* GetNovelSeriesProfile
* GetNovelSeriesTitle
* GetNovelContent
* GetNovelComments

## Known issues
* Golang struct type is ambiguous, usually due to lack of data support.
* Golang struct lacks some data fields, which may be due to lack of maintenance.
* Login API is not available, because reCAPTCHA could not be verified.
* API that can only be used after login.

# About issue
If possible, please provide the source code and API response data that caused the error. Although the lack of such information will make it difficult to locate the error, you are free to decide whether to provide it.

# Contribute code
The project has only recently been opened to the public, so I don't know how to handle requests like this.
~~You may not like my programming style.~~

[^1]: Similar to GIF
