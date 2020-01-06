final static def genMd5(String text, String charset = "UTF-8") {
    return convertToHex(MessageDigest.getInstance("MD5").digest(text.getBytes(charset)))
}

final static def genSha256V1(String key, String text, String charset = "UTF-8") {
    Mac instance = Mac.getInstance("HmacSHA256")
    SecretKeySpec secretKeySpec = new SecretKeySpec(key.getBytes(charset), "HmacSHA256")
    instance.init(secretKeySpec)
    return convertToHex(instance.doFinal(text.getBytes(charset)))
}

final static def genSha256V2(String key, String text, String charset = "UTF-8") {
    Mac instance = Mac.getInstance("HmacSHA256")
    SecretKeySpec secretKeySpec = new SecretKeySpec(key.getBytes(charset), "HmacSHA256")
    instance.init(secretKeySpec)
    return Hex.encodeHexString(instance.doFinal(text.getBytes(charset)))
}

final static def convertToHex(byte[] data) {
    StringBuilder sb = new StringBuilder()

    for (int i = 0; i < data.length; ++i) {
        String hex = Integer.toHexString(data[i] & 0xFF)

        if (hex.length() < 2) {
            sb.append("0")
        }

        sb.append(hex)
    }

    return sb.toString()
}

public static void main(String[] args) {
    println calculateMd5("451279738422597632")
    println genMd5("451279738422597632")
    println genSha256V1(")#@\$%^&*!(", "451279738422597632")
    println genSha256V2(")#@\$%^&*!(", "451279738422597632")
    println HmacUtils.hmacSha256Hex(")#@\$%^&*!(", "451279738422597632")
}
