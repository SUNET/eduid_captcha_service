# EduID_captcha_service
### Compile
`$ make`
creates a amd64 linux binary

### Run
`$ ./bin/eduid_captcha_service`

## Fetch captcha
### Request
JSON format

#### Request:
| Value              | Default | Type   | Multiplicity |
|--------------------|---------|--------|--------------|
| picture_format     | jpeg    | string | 0..1         |
| curve_number       | 2       | int    | 0..1         |
| text_length        | 4       | int    | 0..1         |
| size.width         | 400     | int    | 0..1         |
| size.height        | 200     | int    | 0..1         |
| noise              | 4.0     | float  | 0..1         |
| background_color.R | 0       | int    | 0..1         |
| background_color.G | 0       | int    | 0..1         |
| background_color.B | 0       | int    | 0..1         |
| background_color.A | 0       | int    | 0..1         |
| font_dpi           | 72.0    | float  | 0..1         |
| char_preset        | ABCDEFGHIJKLMNOPQRSTUVWXYZ <br> abcdefghijklmnopqrstuvwxyz <br> 0123456789 | string | 0..1 |


### Reply
JSON format

| Value | Type  | Multiplicity |
|-------|-------|--------------|
| data  | data  | 0..1         |
| error | error | 0..1         |

#### data:
| Value                  | Type    | Description | Multiplicity            |
|------------------------|---------|--------------|------------------------|
| text                   | string  | text representation of captcha | 1..1 |
| picture_base64_encoded | string  | base64 encoded picture         | 1..1 |
| meta                   | Request |                                | 1..1 |


#### error:
| Value   | Type      | Multiplicity |
|---------|-----------|--------------|
| title   | string    | 1..1         |
| details | interface | 1..*         |
