import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Scanner;

public class Main {
	public static void main(String[] args) {

		Pizzeria pizzeria = new Pizzeria();
		CoffeeShop coffeeShop = new CoffeeShop();
		DumplingsShop dumplingsShop = new DumplingsShop();

		Menu mainMenu = new Menu("Menu");
		mainMenu.add(pizzeria.getMenu());
		mainMenu.add(coffeeShop.getMenu());
		mainMenu.add(dumplingsShop.getMenu());

		mainMenu.print();

		ArrayList<String> order = new ArrayList<>();

		System.out.println("type 'done' to quit");
		System.out.print("product to order: ");
		Scanner scanner = new Scanner(System.in);
		String input = scanner.nextLine();
		while (!input.equals("done")) {
			String item = search(mainMenu, input);
			if (item == null) {
				System.out.println("there is no " + input + " in menu");
			} else {
				order.add(item);
				printOrder(order);
			}

			System.out.print("product to order: ");
			input = scanner.nextLine();
		}
		scanner.close();

	}

	private static void printOrder(List<String> order) {
		System.out.println("ORDER:");
		for (String item : order) {
			System.out.println(item);
		}
	}

	private static String search(Menu menu, String input) {
		Iterator<IMenuItem> iterator = menu.iterator();
		while (iterator.hasNext()) {
			IMenuItem item = iterator.next();
			if (item instanceof Menu) {
				String foundName = search((Menu) item, input);
				if (foundName != null) {
					return foundName;
				}
			} else {
				if (item.getName().equalsIgnoreCase(input)) {
					return item.getName();
				}
			}
		}

		return null;
	}
}
