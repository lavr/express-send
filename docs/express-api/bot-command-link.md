# Ссылки на ботов с отправкой команды

iOS 3.17Android 3.17Web 3.17

Для перехода в чат с ботом и отправки ему команды предусмотрены специальные ссылки (deeplink).

## Формат ссылки

Ссылка имеет следующий формат:

```
https://xlnk.ms/open/bot/?ets_id=&body=&command=
```

| Параметр | Тип данных | Описание |
| --- | --- | --- |
| huid | string | Huid бота |
| ets_id | string | ID ETS сервера,необязательный |
| body | string | Текст сообщения. Выведется в бабле сообщения,необязательный |
| command | string | Текст команды. Не доступен пользователю, доступен боту,необязательный |

********
> **Note**
> Примечание
> Поля
> body
> и
> command
> передаются только совместно, то есть в ссылке указываются оба поля или не указывается ни одно из них.

## Генерация QR с командой боту


#### HUID бота




#### Команда боту


        Текст сообщения






        Команда






#### ETS




#### Хост сервера ссылок




        [Сгенерировать](javascript:submit())







            []()




        const UUID_REGEX = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-5][0-9a-f]{3}-[089ab][0-9a-f]{3}-[0-9a-f]{12}$/i;

        function generateQr() {
            var huid = $("#huid").val();
            var ets = $("#ets").val();
            var body = $("#body").val();
            var command = $("#command").val();
            var host = $("#host").val();

            var url = `https://${host}/open/bot/${huid}`
            var query = [{ ets_id: ets }, { body }, { command }]
                .map(val => Object.entries(val))
                .filter(([[key, val]]) => !!val)
                .map(([[key, val]]) => `${key}=${encodeURIComponent(val)}`)
                .join("&");

            if (query) {
                url = `${url}?${query}`;
            }

            $("#qr-url").attr("href", url).text(url);

            var modal = $("#qr-wrapper .qr-modal");

            new QRCode(modal.find(".qr-modal__img").empty()[0], {
                text: url,
                width: 1024,
                height: 1024,
                colorDark : "#000000",
                colorLight : "#ffffff",
                correctLevel : QRCode.CorrectLevel.H
            });

            modal.show();
        }

        function closeModal() {
            $("#qr-wrapper .qr-modal").hide();
        }

        function hideErrors() {
            $("#qr-wrapper input").removeClass("qr-input--error");
            $("#qr-wrapper .qr-error").hide();
            closeModal();

            return true;
        }

        function showAndReturnError(id, text) {
            $(`#${id}`).addClass("qr-input--error");
            $(`#${id}-error`).text(text).show();

            return false;
        }

        function checkHuid() {
            var huid = $("#huid").val();
            if (!huid) return showAndReturnError("huid", "Обязательное поле");

            var success = UUID_REGEX.test(huid);
            if (!success) return showAndReturnError("huid", "Неверный формат");

            return true;
        }

        function checkCommand() {
            var body = $("#body").val();
            var command = $("#command").val();

            if (!body && !!command) return showAndReturnError("body", "Необходимо указать оба поля или оставить пустыми");
            if (!!body && !command) return showAndReturnError("command", "Необходимо указать оба поля или оставить пустыми");

            if (!!body && body.length > 16) return showAndReturnError("body", "Максимум 16 символов");
            if (!!command && command.length > 16) return showAndReturnError("command", "Максимум 16 символов");

            return true;
        }

        function checkEts() {
            var ets = $("#ets").val();

            if (!!ets && !UUID_REGEX.test(ets)) return showAndReturnError("ets", "Неверный формат");

            return true;
        }

        function checkHost() {
            var host = $("#host").val();
            if (!host) return showAndReturnError("host", "Обязательное поле");

            return true;
        }

        function submit() {
            hideErrors() && checkHuid() && checkCommand() && checkEts() && checkHost() && generateQr();
        }

        $(document).ready(hideErrors);
