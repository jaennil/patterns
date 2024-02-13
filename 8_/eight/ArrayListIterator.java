package eight;
import java.util.List;

class ArrayListIterator implements MenuIterator {
    private List<MenuItem> items;
    private int position = 0;

    ArrayListIterator(List<MenuItem> items) {
        this.items = items;
    }

    @Override
    public boolean hasNext() {
        return position < items.size();
    }

    @Override
    public MenuItem next() {
        if (hasNext()) {
            MenuItem menuItem = items.get(position);
            position++;
            return menuItem;
        }
        return null;
    }
}
