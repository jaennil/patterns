import java.util.ArrayList;

public class Pizzeria extends Restaurant {
	public Pizzeria() {
		menu = new Menu("Pizzeria Menu");
		menu.add(new MenuItem("pepperony"));
		menu.add(new MenuItem("veggy"));
		menu.add(new MenuItem("cheese"));
		menu.add(new MenuItem("mashroom"));
	}
}

