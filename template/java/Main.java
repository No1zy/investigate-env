import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

public class Main {
    public static void main(String[] args) {

        String strUrl = "{{ .URL }}"; //"https://example.com＃@bing.com";
        try {

            URL url = new URL(strUrl);
            System.out.println("protocol = " + url.getProtocol());
            System.out.println("UserInfo = " + url.getUserInfo());
            System.out.println("authority = " + url.getAuthority());
            System.out.println("host = " + url.getHost());
            System.out.println("port = " + url.getPort());
            System.out.println("path = " + url.getPath());
            System.out.println("query = " + url.getQuery());
            System.out.println("filename = " + url.getFile());
            System.out.println("ref = " + url.getRef());
        } catch (MalformedURLException e) {
            System.out.println(e);
        }

        HttpURLConnection urlConn = null;
        InputStream in = null;
        BufferedReader reader = null;


        try {
            //接続するURLを指定する
            URL url = new URL(strUrl);

            //コネクションを取得する
            urlConn = (HttpURLConnection) url.openConnection();
            urlConn.setInstanceFollowRedirects(false);

            urlConn.setRequestMethod("GET");
//			urlConn.setRequestMethod("POST");

            urlConn.connect();

            int status = urlConn.getResponseCode();

            System.out.println(urlConn.getURL());
;
            System.out.println("HTTP status:" + status);

            if (status == HttpURLConnection.HTTP_OK) {

                in = urlConn.getInputStream();

                reader = new BufferedReader(new InputStreamReader(in));

                StringBuilder output = new StringBuilder();
                String line;

                while ((line = reader.readLine()) != null) {
                    output.append(line);
                }
                System.out.println(output.toString());
            }
        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            try {
                if (reader != null) {
                    reader.close();
                }
                if (urlConn != null) {
                    urlConn.disconnect();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
