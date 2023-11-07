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
        double scaleFactor = 0.5;
        PaintWindow paintWindow = new PaintWindow(3000, 2400);
        int newWidth = (int) (paintWindow.getWidth() * scaleFactor);
        int newHeight = (int) (paintWindow.getHeight() * scaleFactor);
        paintWindow.setSize(newWidth, newHeight);
        //paintWindow.setCanvasPreferredSize(newWidth, newHeight);
        paintWindow.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
    }
}