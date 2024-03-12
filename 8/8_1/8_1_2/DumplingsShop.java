import java.util.Hashtable;

public class DumplingsShop extends Restaurant<Hashtable<Integer, String>> {
	public DumplingsShop() {
		menu = new Hashtable<Integer, String>();
		menu.put(0, "dumpling 1");
		menu.put(1, "dumpling 2");
		menu.put(2, "dumpling 3");
		menu.put(3, "dumpling 4");
	}
}

