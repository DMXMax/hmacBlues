import java.security.MessageDigest;
import java.nio.charset.StandardCharsets;
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.io.ByteArrayOutputStream;
import java.io.IOException;



class Playground {
    public static void main(String[] args) throws IOException {
        System.out.println("Hello Java");

        System.out.println(bytesToHex("Hello Java".getBytes(StandardCharsets.UTF_8)));
        SecureRandom random = new SecureRandom();

        ByteArrayOutputStream baos = new ByteArrayOutputStream();


        String message = "My Secret Data";
        String salt = "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e";
        //random.nextBytes(salt);
        baos.write(salt.getBytes(StandardCharsets.UTF_8));
        baos.write(message.getBytes(StandardCharsets.UTF_8));

        byte[] encodedhash = {};

        try {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            encodedhash = digest.digest(
                baos.toByteArray());
        } catch (NoSuchAlgorithmException nsae) {
            System.out.println(nsae.toString());

        }

        System.out.println(bytesToHex(encodedhash));
    }


    private static String bytesToHex(byte[] hash) {
        StringBuffer hexString = new StringBuffer();
        for (int i = 0; i < hash.length; i++) {
            String hex = Integer.toHexString(0xff & hash[i]);
            if (hex.length() == 1) 
                hexString.append('0');
            hexString.append(hex);
            hexString.append(' ');
        }
        return hexString.toString();
    }
}
