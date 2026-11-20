# pau

Pau - это минималистичный язык разметки который компилируеться в html

в качестве примера создадим файл `mysite.pau`

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
после чего можно открыть html файл в браузере

## Теги

Pau потдеживает такие теги:
- заголовок - `<h2>`
	- `<h2> TITLE`

- гиперссылки `<a>`
	- `<a> "index.html" | HOME`
	- `<a> "index.html"`

- тег `<ul></ul>` - добавляет табуляцию, покачто в нем можно размещать только гиперссылки 
```txt
<ul>
<a> "index.html" | HOME
<a> "index.html"
</ul>
```

- `<hr>` - добавляет полосу

- любая строка без тега принимается как текст, кроме первой - она для заголовка сайта
- пример:
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

### CLI `./pau -help`

```shell
=== help ===
=== run
        run 'pathToFile.pau'
        run .

```
Скомпилировать все файлы .pau `./pau run .`, дейсвует только в локальной директории.

## License
This project is dedicated to the public domain under the Creative Commons Zero v1.0 Universal license. You can copy, modify, and distribute the code without asking for permission. See the [LICENSE](LICENSE) file for details.
