import 'package:dio/dio.dart';
import 'package:frontend/core/config/config.dart';

class TransactionService {
  final Dio _dio;
  final String _apiUrl;

  TransactionService({Dio? dio})
    : _dio = dio ?? Dio(BaseOptions(baseUrl: Config.apiUrl)),
      _apiUrl = Config.apiUrl;

  String get apiUrl => _apiUrl;

  /// Buscar uma transação pelo ID e moeda
  Future<Map<String, dynamic>> fetchTransaction(
    String id,
    String currency,
  ) async {
    final response = await _dio.get(
      '/transactions/$id',
      queryParameters: {'currency': currency},
    );
    return response.data as Map<String, dynamic>;
  }

  /// Buscar a lista de moedas disponíveis
  Future<List<String>> fetchCurrencies() async {
    final response = await _dio.get('/currencies');
    final List<dynamic> data = response.data;
    return data.map((e) => e.toString()).toList();
  }

  /// Criar uma nova transação
  Future<void> createTransaction(Map<String, dynamic> payload) async {
    await _dio.post('/transactions', data: payload);
  }

  /// Buscar as últimas transações, com limite definido por query param
  Future<List<Map<String, dynamic>>> fetchLatestTransactions({
    int limit = 5,
  }) async {
    final response = await _dio.get(
      '/transactions/latest',
      queryParameters: {'limit': limit},
    );
    final List<dynamic> data = response.data;
    return data.cast<Map<String, dynamic>>();
  }
}
