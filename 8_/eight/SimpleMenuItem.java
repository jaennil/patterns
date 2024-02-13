package eight;
class SimpleMenuItem extends MenuItem {
    SimpleMenuItem(String name, double price) {
        super(name, price);
    }

    @Override
    void print() {
        System.out.println(name + " - $" + price);
    }
}
