## ohgibone v0.22.0

[![GoDoc](https://godoc.org/github.com/hico-horiuchi/ohgibone/sensu?status.svg)](https://godoc.org/github.com/hico-horiuchi/ohgibone/sensu) [![Circle CI](https://circleci.com/gh/hico-horiuchi/ohgibone.svg?style=shield)](https://circleci.com/gh/hico-horiuchi/ohgibone) ![Coverage](https://img.shields.io/badge/coverage-76.8%25-lightgray.svg)

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Sensu](http://sensuapp.org/) >= 0.19

#### Documents

  - [github.com/hico-horiuchi/ohgibone/sensu](http://godoc.org/github.com/hico-horiuchi/ohgibone/sensu) 

#### Installation

    $ go get github.com/hico-horiuchi/ohgibone/sensu

#### Coverage

| NAMESPACE                    | METHOD   | IMPRIMENTATION |
|:-----------------------------|:---------|:---------------|
| `/clients`                   | `GET`    | DONE           |
| `/clients`                   | `POST`   | DONE           |
| `/clients/:client`           | `GET`    | DONE           |
| `/clients/:client`           | `DELETE` | DONE           |
| `/clients/:client/history`   | `GET`    | -> `history`   |
| `/checks`                    | `GET`    | DONE           |
| `/checks/:check`             | `GET`    | DONE           |
| `/request`                   | `POST`   | DONE           |
| `/events`                    | `GET`    | DONE           |
| `/events/:client`            | `GET`    | DONE           |
| `/events/:client/:check`     | `GET`    | DONE           |
| `/events/:client/:check`     | `DELETE` | DONE           |
| `/resolve`                   | `POST`   | DONE           |
| `/results`                   | `GET`    | DONE           |
| `/results/:client`           | `GET`    | DONE           |
| `/results/:client/:check`    | `GET`    | DONE           |
| `/aggregates`                | `GET`    | DONE           |
| `/aggregates/:check`         | `GET`    | DONE           |
| `/aggregates/:check`         | `DELETE` | DONE           |
| `/aggregates/:check/:issued` | `GET`    | DONE           |
| `/stashes`                   | `GET`    | DONE           |
| `/stashes`                   | `POST`   | DONE           |
| `/stashes/:path`             | `POST`   | DONE           |
| `/stashes/:path`             | `GET`    | DONE           |
| `/stashes/:path`             | `DELETE` | DONE           |
| `/health`                    | `GET`    | DONE           |
| `/info`                      | `GET`    | DONE           |

#### License

ohgibone is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/ohgibone/master/LICENSE).
