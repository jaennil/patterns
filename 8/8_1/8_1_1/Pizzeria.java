import java.util.ArrayList;

public class Pizzeria extends Restaurant<ArrayList<String>> {
	public Pizzeria() {
		menu = new ArrayList<>();
		menu.add("pepperony");
		menu.add("veggy");
		menu.add("cheese");
		menu.add("mashroom");
	}
}

