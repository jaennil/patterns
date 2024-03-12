import java.util.Vector;

public class CoffeeShop extends Restaurant {
	public CoffeeShop() {
		menu = new Menu("CoffeeShop Menu");
		menu.add(new MenuItem("espresso"));
		menu.add(new MenuItem("americano"));
		menu.add(new MenuItem("dark roast"));
		menu.add(new MenuItem("decaf"));
	}
}

