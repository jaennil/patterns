import java.util.Scanner;
import java.util.ArrayList;
import java.util.Vector;
import java.util.Hashtable;

public class Main {
	public static void main(String[] args) {

		System.out.println("1. Pizzeria");
		System.out.println("2. CoffeeShop");
		System.out.println("3. DumplingsShop");
		Scanner scanner = new Scanner(System.in);
		String choice = scanner.nextLine();
		scanner.close();

		Restaurant<?> restaurant;
		MenuIterator menuIterator;
		if (choice.equals("1")) {
			restaurant = new Pizzeria();
			menuIterator = new PizzeriaMenuIterator((ArrayList<String>) restaurant.getMenu());
		} else if (choice.equals("2")) {
			restaurant = new CoffeeShop();
			menuIterator = new CoffeeShopMenuIterator((Vector<String>) restaurant.getMenu());
		} else {
			restaurant = new DumplingsShop();
			menuIterator = new DumplingsShopMenuIterator((Hashtable<Integer, String>) restaurant.getMenu());
		}

		System.out.println("print menu");
		while (menuIterator.hasNext()) {
			System.out.println(menuIterator.next());
		}

		menuIterator.reset();

		System.out.println("delete all");
		while (menuIterator.hasNext()) {
			while (menuIterator.hasNext()) {
				System.out.println(menuIterator.next());
			}
			menuIterator.removeLast();
			menuIterator.reset();
		}

		menuIterator.reset();

		System.out.println("after delete");
		while (menuIterator.hasNext()) {
			System.out.println(menuIterator.next());
		}
	}
}
