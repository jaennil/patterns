package eight;
import java.util.ArrayList;
import java.util.List;

class Order {
    List<MenuItem> items = new ArrayList<>();

    void addItem(MenuItem item) {
        items.add(item);
    }

    void printOrder() {
        double totalPrice = 0;
        for (MenuItem item : items) {
            item.print();
            totalPrice += item.price;
        }
        System.out.println("Total: $" + totalPrice);
    }
}
