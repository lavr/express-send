# Платформа BotX: HTTP(s) BotX API

## Ограничения API

### Длина запроса

Общая длина запроса (Content-Length) не должна превышать 512 МБ.
Суммарный вес JSON-полей запроса (исключая вес файла) не должен превышать 1 МБ.

### Размер передаваемого файла

Размер передаваемого файла не должен превышать 512 МБ, что составляет ~536 МБ в base64-представлении.

### Расширения файлов

При валидации файла расширение файла определяется на основе переданного `mimetype` и таблицы соответствия (см. таблицу соответствия `mimetype` и расширений):

- если `mimetype` присутствует в таблице, то файл сохраняется с соответствующим расширением и переданным `mimetype`;

- если `mimetype` отсутствует в таблице, то файл сохраняется, как неопределенный бинарный файл с `mimetype` "application/octet-stream".

**Таблица соответствия mimetype и расширений**

```
%{
    "application/epub+zip" => ["epub"],
    "application/gzip" => ["gz"],
    "application/java-archive" => ["jar"],
    "application/javascript" => ["js"],
    "application/json" => ["json"],
    "application/json-patch+json" => ["json-patch"],
    "application/ld+json" => ["jsonld"],
    "application/manifest+json" => ["webmanifest"],
    "application/msword" => ["doc"],
    "application/octet-stream" => ["bin"],
    "application/ogg" => ["ogx"],
    "application/pdf" => ["pdf"],
    "application/postscript" => ["ps", "eps", "ai"],
    "application/rtf" => ["rtf"],
    "application/vnd.amazon.ebook" => ["azw"],
    "application/vnd.api+json" => ["json-api"],
    "application/vnd.apple.installer+xml" => ["mpkg"],
    "application/vnd.mozilla.xul+xml" => ["xul"],
    "application/vnd.ms-excel" => ["xls"],
    "application/vnd.ms-fontobject" => ["eot"],
    "application/vnd.ms-powerpoint" => ["ppt"],
    "application/vnd.oasis.opendocument.presentation" => ["odp"],
    "application/vnd.oasis.opendocument.spreadsheet" => ["ods"],
    "application/vnd.oasis.opendocument.text" => ["odt"],
    "application/vnd.openxmlformats-officedocument.presentationml.presentation" => ["pptx"],
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" => ["xlsx"],
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document" => ["docx"],
    "application/vnd.rar" => ["rar"],
    "application/vnd.visio" => ["vsd"],
    "application/x-7z-compressed" => ["7z"],
    "application/x-abiword" => ["abw"],
    "application/x-bzip" => ["bz"],
    "application/x-bzip2" => ["bz2"],
    "application/x-cdf" => ["cda"],
    "application/x-csh" => ["csh"],
    "application/x-freearc" => ["arc"],
    "application/x-httpd-php" => ["php"],
    "application/x-msaccess" => ["mdb"],
    "application/x-sh" => ["sh"],
    "application/x-shockwave-flash" => ["swf"],
    "application/x-tar" => ["tar"],
    "application/xhtml+xml" => ["xhtml"],
    "application/xml" => ["xml"],
    "application/wasm" => ["wasm"],
    "application/zip" => ["zip"],
    "audio/3gpp" => ["3gp"],
    "audio/3gpp2" => ["3g2"],
    "audio/aac" => ["aac"],
    "audio/midi" => ["mid", "midi"],
    "audio/mpeg" => ["mp3"],
    "audio/ogg" => ["oga"],
    "audio/opus" => ["opus"],
    "audio/wav" => ["wav"],
    "audio/webm" => ["weba"],
    "font/otf" => ["otf"],
    "font/ttf" => ["ttf"],
    "font/woff" => ["woff"],
    "font/woff2" => ["woff2"],
    "image/avif" => ["avif"],
    "image/bmp" => ["bmp"],
    "image/gif" => ["gif"],
    "image/jpeg" => ["jpg", "jpeg"],
    "image/png" => ["png"],
    "image/svg+xml" => ["svg", "svgz"],
    "image/tiff" => ["tiff", "tif"],
    "image/vnd.microsoft.icon" => ["ico"],
    "image/webp" => ["webp"],
    "text/calendar" => ["ics"],
    "text/css" => ["css"],
    "text/csv" => ["csv"],
    "text/html" => ["html", "htm"],
    "text/javascript" => ["js", "mjs"],
    "text/plain" => ["txt", "text"],
    "text/xml" => ["xml"],
    "video/3gpp" => ["3gp"],
    "video/3gpp2" => ["3g2"],
    "video/quicktime" => ["mov"],
    "video/mp2t" => ["ts"],
    "video/mp4" => ["mp4"],
    "video/mpeg" => ["mpeg", "mpg"],
    "video/ogg" => ["ogv"],
    "video/webm" => ["webm"],
    "video/x-msvideo" => ["avi"],
    "video/x-ms-wmv" => ["wmv"]
 }
```

## Версия сервиса

Версию сервиса можно получить по следующему URL:

```
GET /system/botx/version
```

Пример ответа:

```
{
  "version": "3.21.0"
}
```

Версия отдается только в релизных сборках. В сборках от master-ветки и от кастомных тегов версия будет некорректная.

## Методы BotX API

- [Bots API](bots-api/)

- [Notifications API](notifications-api/)

- [Events API](events-api/)

- [Chats API](chats-api/)

- [Users API](users-api/)

- [Files API](files-api/)

- [SmartApps API](../../../../smartapps/developer-guide/smartapp-api/)

- [Stickers API](stickers-api/)

- [OpenID API](openid-api/)

- [Metrics API](metrics-api/)

- [VoEx API](voex-api/)
