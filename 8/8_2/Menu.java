import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class Menu implements IMenuItem {
	private String name;
	private List<IMenuItem> items;

	public Menu(String name) {
		this.name = name;
		items = new ArrayList<>();
	}

	public void add(IMenuItem menuItem) {
        items.add(menuItem);
    }

    public void remove(IMenuItem menuItem) {
        items.remove(menuItem);
    }

    public void print() {
        System.out.println(name);
        for (IMenuItem item : items) {
            item.print();
        }
    }

	@Override
	public String getName() {
		return name;
	}

	public Iterator<IMenuItem> iterator() {
		return new MenuIterator(items);
	}
}
