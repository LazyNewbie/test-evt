<?php
declare(strict_types=1);

$phpFiles = $argv[1] ?? null;
$gbpMetaDir = $argv[2] ?? null;

if(is_null($phpFiles)){
    panic('php files base dir need to be provided');
}

if(is_null($gbpMetaDir)){
    panic('GPBMetadata base dir name need to be provided (not a directory path)');
}

$dirIterator = new RecursiveIteratorIterator(new RecursiveDirectoryIterator($phpFiles, FilesystemIterator::CURRENT_AS_FILEINFO|FilesystemIterator::SKIP_DOTS));
$files = array();

$filesModified = [];

/** @var SplFileInfo $file */
foreach ($dirIterator as $file) {
    if ($file->isDir()){
        continue;
    }

    $content = file_get_contents($file->getPathname());

    if(strpos($content, ' const ') !== false ){
        file_put_contents($file->getPathname(), makeEnum($content));
        $filesModified[] = $file->getPathname();
    }
}

info('Files modified ('.count($filesModified)."):\n".implode("\n", $filesModified));


function makeEnum(string $content): string
{

    preg_match_all('~const (\S+) = (\d+);~i', $content, $matches);
    $cases = [];

    foreach ($matches[0] as $i=>$line){
        $cases[] = "case {$matches[1][$i]} = {$matches[2][$i]};";
    }
    $cases = implode("\n    ", $cases);

    preg_match('~namespace (\S+);~', $content, $match);
    $namespace = $match[1];

    preg_match('~class (\S+)~', $content, $match);
    $className = $match[1];


    $headerComments = substr($content, 0, strpos($content, 'namespace'));


    return <<<EOF
$headerComments
namespace $namespace;

enum $className: int
{
    $cases
}
EOF;

}

function panic(string $message): void
{
    echo "[ERROR]: $message\n";
    exit(1);
}

function info(string $message): void
{
    echo "[INFO]: $message\n";
}