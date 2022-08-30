# Masterarbeit Wenkemann David (FH Erfurt)
***

## Table of Contents
1. [General Info](#general-info)
2. [Technologies](#technologies)
3. [Installation](#installation)
4. [Collaboration](#collaboration)
5. [FAQs](#faqs)

### General Info
***
Code zur Masterarbeit von David Wenkemann

Es wurde ein einfaches Beispiel im Szenario eines Getränkehändlers implementiert.
Es gibt vier Module: Basedata, Reporting, Store, Warehouse

Die Applikation wurde in zwei verschiedenen Varianten implementiert, welche in den jeweiligen Unterordnern liegen.
* Monolith mit synchronem Kommunikationspattern
* Monolith mit asynchronem Kommunikationspattern

Wenn Beispieldaten genutzt werden sollen, muss in der main.go die Funktion database.SpinupDB() gestartet werden.


## Technologies
***
Eine Liste der im Projekt genutzten Technologien:
* [Golang](https://go.dev/): Version 1.19 
* [BubbleTea](github.com/charmbracelet/bubbletea): Version 0.22.1

## Installation
***
Um die Applikation herunterzuladen und zu installieren muss folgendes getan werden:
```
$ git clone https://github.com/DavidWenkemann/Masterarbeit.git
$ cd ../path/to/the/file
$ make compile
$ go run
```

## FAQs
***
Liste der meistgestellten Fragen:
1. **Wird die Applikation weiterentwickelt?**
Nein. Dies ist eine abgeschlossene Masterarbeit 
