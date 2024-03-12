import java.util.Hashtable;

public class DumplingsShop extends Restaurant {
	public DumplingsShop() {
		menu = new Menu("DumplingsShop Menu");
		menu.add(new MenuItem("dumpling 1"));
		menu.add(new MenuItem("dumpling 2"));
		menu.add(new MenuItem("dumpling 3"));
		menu.add(new MenuItem("dumpling 4"));

		Menu wineMenu = new Menu("DumplingsShop Wine Menu");
		wineMenu.add(new MenuItem("wine 1"));
		wineMenu.add(new MenuItem("wine 2"));
		wineMenu.add(new MenuItem("wine 3"));
		menu.add(wineMenu);
	}
}

