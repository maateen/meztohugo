
# MezToHugo

MezToHugo is a modern tool to migrate from [Mezzanine](http://mezzanine.jupo.org/) to [Hugo](https://gohugo.io/), the world’s fastest framework for building websites. Cuurently it migrates posts only.

## Usage

1. Download the binary file from releases.
2. Create `config.yml` file in the working directoty where you will run the too.
3. Paste the following things into config file.
```yml
database:
  type: "mysql"
  hostname: "127.0.0.1"
  port: "3306"
  dbname: "en_maateen"
  username: "root"
  password: "password"
hugo:
  contentDir: "post"
  imageDir: "img"
  defaultCoverImage: "/img/cover.jpg"
```
4. Now run the tool from terminal.
```
$ chmod +x meztohugo
$ ./meztohugo
```

## To-Do

Please contribute and complete the project if you can.

- [ ] Add support to export page.