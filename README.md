![Thumbnail](./images/dotcanvas_en.png)  
# dot Canvas
> **r/place** meterialized with a Discord BOT

You can also read it in the following languages:  
**English** | [한국어](./docs/korean.md)  
* * *
## User's Guide
If you're filling in the color, you can fill it in again after a minute.

### Fill
The `fill` command allows you to fill in the selected color in the coordinates you want.
### Fill Hex
You can fill the coordinates you want with the color you entered through the `fillhex` command. 
However, the color you are entering should be a color code method (#000000).
### Canvas
You can check the canvas filled up to now through the `canvas` command. 
The canvas is initialized at midnight on the 1st of every month.
* * *
## CANVAS
![canvas]](https://dotcanvas.neoration.me/image)
If the bot server is not working, it may not be visible.  
* * *
## Getting Started
This project works on Go 1.18 and above.  

### To set Env
In order to run the bot, we need an `.env` file containing the information below at the top of the repository.  
Please check the file `example.env` for more information.
```
TOKEN=Write-Your-Discord-Bot-Token
CLIENT_ID=Write-Your-Discord-Bot-ID
MONGO_URI=Write-Your-MongoDB-URI
DATABASE_NAME=Write-Your-MongoDB-Database-Name
IMAGE_URL=Write-Canvas-ImageURL
```  

### To check the port
This project is a form of showing pictures of canvas directly through the web.  
Make sure port 80 is available and working properly.  
* * *
## Contributing
### Code Contribution
If you notice incorrectly written or inefficient code, please fork out this repository and create a Pull Request with your own code!  

### Translation Contribution
You can add another language by translating to the path `src/locales`.  
In this case, please create a Pull Request after you fork the repository yourself, or send the file through the [creator's Discord](https://discord.com/users/726534821572116512).

### Support server costs
You can support the cost of the server [here](https://toss.me/neorate).  