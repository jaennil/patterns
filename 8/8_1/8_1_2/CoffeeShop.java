import java.util.Vector;

public class CoffeeShop extends Restaurant<Vector<String>> {
	public CoffeeShop() {
		menu = new Vector<String>();
		menu.add("espresso");
		menu.add("americano");
		menu.add("dark roast");
		menu.add("decaf");
	}
}

