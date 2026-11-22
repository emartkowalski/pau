# pau

Pau - это минималистичный язык разметки который компилируеться в html

для начала создадим файл `mysite.pau`
```txt
Personal Site

hello everyone!
i am a human!
check this out...
```

затем выполним компиляцию
```shell
./pau run mysite.pau
```
после чего можно открыть html файл в браузере.

### Documentation

Pau потдеживает такие теги: <br>
Заголовки:
<br> `<h2> TITLE` - компилируеться в `<h2>TITLE</h2>`

Гиперссылки:
<br> `<a> "index.html" | HOME`
<br> `<a> "index.html"`

`<ul></ul>` - добавляет табуляцию, покачто в нем можно размещать только гиперссылки 
```txt
<ul>
<a> "index.html" | HOME
<a> "index.html"
</ul>
```

`<hr>` - добавляет полосу
<br> Любая строка без тега принимается как текст, кроме первой, она для заголовка сайта.
<br> Пример:
```txt
Personal Site

<h2> Navigation
<ul>
<a> "https://en.wikipedia.org/wiki/Internet" | Internet
<a> "https://en.wikipedia.org/wiki/Cactus"   | Cactus
<a> "https://en.wikipedia.org/wiki/Dog"      | Dogs
<a> "https://en.wikipedia.org/wiki/Cat"      | Cats
<a> "https://www.wikipedia.org/"
</ul>

<h2> some text here
the first string
the second string

the end of text...

<hr>
<a> "index.html" | HOME
Personal Site | 2090
```
### args:
- `./pau -version`
- `./pau -help`
- `./pau run .` - Рекурсивный обход вложенных директорий, начиная с текущей

## License
This project is dedicated to the public domain under the Creative Commons Zero v1.0 Universal license.You can copy, modify, and distribute the code without asking for permission.See the [LICENSE](LICENSE) file for details.
