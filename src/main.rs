use gtk::prelude::*;
use gtk::{Application, ApplicationWindow, TextView, MenuBar, MenuItem, Box, ScrolledWindow};
use rfd::FileDialog;
use std::fs;
fn main() {
    let application = Application::new(Some("com.example.advanced_text_editor"), Default::default());
    application.connect_activate(|app| {
        let window = ApplicationWindow::new(app);
        window.set_title("Advanced Text Editor");
        window.set_default_size(800, 600);
        let vbox = Box::new(gtk::Orientation::Vertical, 0);
        let menu_bar = MenuBar::new();
        let file_menu = gtk::Menu::new();
        let file_menu_item = MenuItem::with_label("File");
        let open_item = MenuItem::with_label("Open");
        let save_item = MenuItem::with_label("Save");
        let exit_item = MenuItem::with_label("Exit");
        open_item.connect_activate(move |_| {
            if let Some(path) = FileDialog::new().pick_file() {
                let content = fs::read_to_string(path.path()).unwrap_or_default();
            }
        });
        save_item.connect_activate(move |_| {
            if let Some(path) = FileDialog::new().save_file() {
                fs::write(path.path(), "Your content here").expect("Unable to write file");
            }
        });
        exit_item.connect_activate(|_| {
            gtk::main_quit();
        });
        file_menu.append(&open_item);
        file_menu.append(&save_item);
        file_menu.append(&exit_item);
        file_menu_item.set_submenu(Some(&file_menu));
        menu_bar.append(&file_menu_item);
        let text_view = TextView::new();
        let scrolled_window = ScrolledWindow::new(None, None);
        scrolled_window.add(&text_view);
        vbox.pack_start(&menu_bar, false, false, 0);
        vbox.pack_start(&scrolled_window, true, true, 0);
        window.add(&vbox);
        window.show_all();
    });
    application.run();
}
