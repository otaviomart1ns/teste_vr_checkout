import 'package:hive/hive.dart';

part 'local_transaction.g.dart';

@HiveType(typeId: 0)
class LocalTransaction extends HiveObject {
  @HiveField(0)
  final String id;

  @HiveField(1)
  final String description;

  @HiveField(2)
  final DateTime date;

  @HiveField(3)
  final double amountUsd;

  LocalTransaction({
    required this.id,
    required this.description,
    required this.date,
    required this.amountUsd,
  });

  LocalTransaction copyWith({
    String? id,
    String? description,
    DateTime? date,
    double? amountUsd,
  }) {
    return LocalTransaction(
      id: id ?? this.id,
      description: description ?? this.description,
      date: date ?? this.date,
      amountUsd: amountUsd ?? this.amountUsd,
    );
  }
}
