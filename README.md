# lander - automatische Landing-Page fuer euren Dockerhost

|Maintainer| David Daehne <david.daehne@msg-david.de>|
|---|---|
|Version|0.1|
|Status|~~geplant~~ -> **in Arbeit** -> Evaluation -> Bereit|
|Sprache|go|

## Wie funktioniert's?
Ganz einfach: Lander ist im Grunde ein Webserver, der bei jeder Anfrage auf '/' den docker-daemon kontaktiert und sich eine Liste der laufenden Container besorgt, 
diese entsprechend seiner Konfiguration evaluiert und danach eine index.html zusammenbastelt.

Um dies zu tun, brauch lander an den entsprechenden Containern folgende label:
- lander.enable: Wenn dieses label gefunden wird, beginnt lander weitere labels zu parsen (der wert ist aktuell noch egal)
- lander.group: Gibt die "Gruppe" an, unter der der Container/die Anwendung in der ausgelieferten index.html eingruppiert wird.
- lander.name: Der Wert dieser Variable entspricht dem angezeigten Text in der index.html, der auf diesen Container verlinkt.

## Konfiguration
Lander wird vollstaendig ueber Umgebungsvariablen konfiguriert. Aktuell stehen folgende Optionen zur Auswahl:

- *LANDER_DOCKER*: der vollstaendige Pfad (einschliesslich Protokoll), zum socket des zu verwendenden Docker daemons (bsp. "unix:///var/run/docker.sock" - Standard unter Linux)
- LANDER_TRAEFIK: gibt an, ob lander nach [Traefik](https://traefik.io/) spezifischen labeln an docker-containern suchen und verwenden soll. Moegliche Werte:
    - true: standard
    - false
- LANDER_EXPOSED: gibt an, ob lander nach Containern suchen soll, von denen Port nach "Aussen", also auf Ports des Docker-Hosts, gemapped wurden. Moegliche Werte:
    - true
    - false: standard
- LANDER_LISTEN: gibt die Adresse an, unter der lander auf http-Anfragen reagieren soll. (bsp: 192.168.1.1:9000, wenn diese Variable nicht angegeben wird, wird sie automatisch auf ":8080" gesetzt)
- LANDER_TITLE: gibt den String an, der in der ausgelieferten index.html als Ueberschrift angezeigt werden soll (standard: "LANDER")
- LANDER_HOSTNAME: sollte auf den Hostnamen der Maschine gesetzt sein, auf dem lander laeuft - wird benutzt um die URLs zu generieren. (Achtung: lander *kann* auch ohne diese Variable richtig funktionieren, dies wird aber nicht garantiert!)

## aktueller Stand:
Aktuell funktioniert nur ein kleiner Teil von lander.. Also eigentlich nur das aufspruehren von "ueber traefik exposede" Container. Wir haben zwar auch eine "main_test.go", aber.. naja, diese Tests sind weder 
gut, noch einigermassen gepflegt.

## Mitmachen:
Wenn du mitmachen... tja.. mal sehen. Am Besten du erstellt dir einen eigenen branch und stellst dann einen Merge Request.
