# Platsbanken - med bättre sökfunktioner
**Ladda ner och gör avancerade sökningar på platsbankens databas**

![Platsbanken screenshot](/preview.png)

## Installation (Docker)
``` bash
git clone https://github.com/linx02/platsbanken.git
```
``` bash
cd platsbanken
docker-compose build
docker-compose up
```

## Funktioner

**Positiv sökning**
- Söker i titel och beskrivning efter angivna kommasepareradesökord (Ej känsligt för versaler)
- Exempel: utvecklare, developer

**Negativ sökning**
- Söker i titel och beskrivning och filtrerar bort annonser med angivna sökord
- Exempel: embedded, inbyggda system, c++

**Avancerad sökning**
- Använd logik för att söka
- Exempel: ("20 års erfarenhet" not in description) and ("senior" not in title)

## Lathund för avancerad sökning (experimentell)

**Format:**
- ("värde" logik fält) and/or ("värde" logik fält) - ""
- ('värde' logik fält) and/or ('värde' logik fält) - ''
- "värde" logik fält - ""
- 'värde' logik fält - ''

**Exempel:**
- ("20 års erfarenhet" not in description) and ("senior" not in title)
- ("javascript" in description) and ("junior" in title)
- ("react" in description) or ("nextjs" in description)
- "junior" in title
- "senior" not in title