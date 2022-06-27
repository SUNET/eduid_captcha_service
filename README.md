# EduID_captcha_service
### Compile
`$ make`
creates a amd64 linux binary

### Run
`$ export EDUID_CONFIG_YAML=dev_config.yaml`

`$ ./bin/eduid_captcha_service`

## Fetch captcha
### Request
JSON format

#### Request:
| Value              | Default | Type             | Multiplicity |
|--------------------|---------|------------------|--------------|
| picture_format     | jpeg    | string           | 0..1         |
| curve_number       | 2       | int              | 0..1         |
| text_length        | 4       | int              | 0..1         |
| size               |         | size             | 0..1         |
| noise              | 4.0     | float            | 0..1         |
| background_color   |         | background_color | 0..1         |
| font_dpi           | 72.0    | float            | 0..1         |
| char_preset        | ABCDEFGHIJKLMNOPQRSTUVWXYZ <br> abcdefghijklmnopqrstuvwxyz <br> 0123456789 | string | 0..1 |

#### size
| Value | Default | Type | Multiplicity |
|--------|--------|------|--------------|
| width  | 400    | int  | 1..1         |
| height | 200    | int  | 1..1         |


#### background_color
| Value | Default | Type       | Description  | Multiplicity |
|-------|---------|------------|--------------|--------------|
| r     | 0       | int        | 0-255, red   | 1..1         |
| g     | 0       | int        | 0-255, green | 1..1         |
| b     | 0       | int        | 0-255, blue  | 1..1         |
| a     | 0       | int        | 0-255, alpha | 1..1         |

### Reply
JSON format

| Value | Type  | Multiplicity |
|-------|-------|--------------|
| data  | data  | 0..1         |
| error | error | 0..1         |

#### data:
| Value | Type    | Description                    | Multiplicity |
|-------|---------|--------------------------------|--------------|
| text  | string  | text representation of captcha | 1..1         |
| image | string  | base64 encoded image           | 1..1         |
| meta  | Request |                                | 1..1         |


#### error:
| Value   | Type      | Multiplicity |
|---------|-----------|--------------|
| title   | string    | 1..1         |
| details | interface | 1..*         |
