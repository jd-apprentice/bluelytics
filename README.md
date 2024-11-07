# Bluelytics

![Wallpaper](https://wallpaperaccess.com/full/4482737.png)

Using the API of [Bluelytics](https://api.bluelytics.com.ar/v2/latest) to get the current value of the dollar in ARS (Argentine Peso) for their different ways of exchange.

Simple project to get started with Infisical.

## Requirements

- [Go](https://golang.org/)

Follow this guide to get started:

- [Getting Started](https://infisical.com/docs/documentation/getting-started/api)

## Example

```shell
$ dolar
+-----------------+----------------+
|   💱 CURRENCY   |    💸 VALUE    |
+-----------------+----------------+
| 💵 Oficial      | 💰 995.00 ARS  |
| 💵 Blue         | 💰 1125.00 ARS |
| 💶 Oficial_Euro | 💰 1081.00 ARS |
| 💶 Blue_euro    | 💰 1223.00 ARS |
+-----------------+----------------+
```

## Dependencies

- infisical
- olekukonko/tablewriter
- joho/godotenv