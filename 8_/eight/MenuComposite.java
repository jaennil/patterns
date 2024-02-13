package eight;

import java.util.ArrayList;
import java.util.List;

class MenuComposite extends MenuItem {
    List<MenuItem> items = new ArrayList<>();

    MenuComposite(String name) {
        super(name, 0);
    }

    void addItem(MenuItem item) {
        items.add(item);
    }

    @Override
    void print() {
        System.out.println(name);
        for (MenuItem item : items) {
            item.print();
        }
    }
}