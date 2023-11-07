package edu.cmu.hcii.paint;
import javax.swing.JFrame;

public class Main {
    public static void main(String[] args) {
        javax.swing.SwingUtilities.invokeLater(new Runnable() {
            public void run() {
                createAndShowGUI();
            }
        });
    }

    private static void createAndShowGUI() {
        PaintWindow paintWindow = new PaintWindow(800, 600);
        paintWindow.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
    }
}