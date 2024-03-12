import java.util.Iterator;
import java.util.List;
import java.util.NoSuchElementException;

public class MenuIterator implements Iterator<IMenuItem> {
	private List<IMenuItem> items;
	private int currentIndex;

	public MenuIterator(List<IMenuItem> items) {
		this.items = items;
		this.currentIndex = 0;
	}

	@Override
	public boolean hasNext() {
		return currentIndex < items.size();
	}

	@Override
	public IMenuItem next() {
		if (!hasNext()) {
			throw new NoSuchElementException();
		}
		return items.get(currentIndex++);
	}
}
