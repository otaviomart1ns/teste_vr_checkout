import 'package:flutter_dotenv/flutter_dotenv.dart';

class Config {
  static String get apiUrl => dotenv.env['API_URL'] ?? 'http://localhost:8080';
}
