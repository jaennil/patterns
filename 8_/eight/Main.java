package eight;
public class Main {
    public static void main(String[] args) {
        
        MenuComposite pizzaMenu = new MenuComposite("Pizzeria");
        pizzaMenu.addItem(new SimpleMenuItem("Margherita", 10.99));
        pizzaMenu.addItem(new SimpleMenuItem("Pepperoni", 12.99));

        MenuComposite coffeeMenu = new MenuComposite("Coffe");
        coffeeMenu.addItem(new SimpleMenuItem("Espresso", 2.99));
        coffeeMenu.addItem(new SimpleMenuItem("Latte", 4.99));

        MenuComposite dessertMenu = new MenuComposite("Desserts");
        dessertMenu.addItem(new SimpleMenuItem("Cake", 6.99));
        coffeeMenu.addItem(dessertMenu);

        Order order = new Order();
        order.addItem(pizzaMenu);
        order.addItem(coffeeMenu);
        
        order.printOrder();
    }
}
