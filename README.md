# pau

## desription
суть pau в том что бы компилировать в html но и в тоже время быть масимально ближе к тексту

## usage:
```bash
pau run mysite.pau
// as alternative you can use .p as an extetion
// or any kind of .txt files
```
и тогда оно сгенерирует файл mysite.html

первая строка файла используеться как заголовок страници

все следущие строки обарачиваются в <p></p> включая переводы строк <br>

пример:
```txt
My Site

hello everyone!
i am a human!
check this out...
```

затем компиляция
```bash
pau run mysite.pau
```
затем можете открыть html файл в браузере

## License
This project is dedicated to the public domain under the Creative Commons Zero v1.0 Universal license. You can copy, modify, and distribute the code without asking for permission. See the [LICENSE](LICENSE) file for details.
