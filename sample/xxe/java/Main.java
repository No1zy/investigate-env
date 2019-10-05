import java.io.FileReader;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;

import org.w3c.dom.Document;
import org.xml.sax.InputSource;

public class Main {

    public static void main(String[] args) throws Exception {

        String filename = "{{ .VARIABLE }}";
        DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
        DocumentBuilder builder = factory.newDocumentBuilder();
        Document doc = builder.parse(new InputSource(filename));
        String data = doc.getElementsByTagName("data").item(0).getTextContent();

        System.out.println("data: " + data);
    }
}
